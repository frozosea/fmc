package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	pb "github.com/frozosea/fmc-pb/tracking"
	scheduler "github.com/frozosea/scheduler/pkg"
	"github.com/go-redis/redis/v8"
	"golang_tracking/pkg/cache"
	"golang_tracking/pkg/logging"
	"golang_tracking/pkg/scac"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/cosu"
	"golang_tracking/pkg/tracking/feso"
	"golang_tracking/pkg/tracking/halu"
	"golang_tracking/pkg/tracking/maeu"
	"golang_tracking/pkg/tracking/mscu"
	"golang_tracking/pkg/tracking/oney"
	"golang_tracking/pkg/tracking/reel"
	"golang_tracking/pkg/tracking/sitc"
	"golang_tracking/pkg/tracking/sklu"
	"golang_tracking/pkg/tracking/util/datetime"
	"golang_tracking/pkg/tracking/util/requests"
	"golang_tracking/pkg/tracking/util/scac_accessory"
	"golang_tracking/pkg/tracking/util/sitc/captcha_resolver"
	"golang_tracking/pkg/tracking/util/sitc/login_provider"
	"golang_tracking/pkg/tracking/zhgu"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

type Builder struct {
	server               *grpc.Server
	variables            *EnvVariables
	db                   *sql.DB
	sitcStore            *login_provider.Store
	captchaSolver        captcha_resolver.ICaptcha
	unclocodesRepository sklu.IRepository
	cache                cache.ICache
	loginProvider        *login_provider.TaskManager

	cosuContainerTracker         *cosu.ContainerTracker
	fesoContainerTracker         *feso.ContainerTracker
	haluContainerTracker         *halu.ContainerTracker
	maeuContainerTracker         *maeu.ContainerTracker
	mscuContainerTracker         *mscu.ContainerTracker
	oneyContainerTracker         *oney.ContainerTracker
	sitcContainerTracker         *sitc.ContainerTracker
	skluContainerTracker         *sklu.ContainerTracker
	scacRepository               scac_accessory.IRepository
	containerTrackingMainTracker *tracking.ContainerTracker
	containerTrackingService     *tracking.ContainerTrackingService
	containerTrackingGrpcService *tracking.ContainerTrackingGrpc

	fesoBillTracker     *feso.BillTracker
	haluBillTracker     *halu.BillTracker
	sitcBillTracker     *sitc.BillTracker
	skluBillTracker     *sklu.BillTracker
	zhguBillTracker     *zhgu.BillTracker
	reelBillTracking    *reel.BillTracker
	billMainTracker     *tracking.BillTracker
	billTrackingService *tracking.BillTrackingService
	billTrackingGrpc    *tracking.BillNumberTrackingGrpc

	scacService     *scac.Service
	scacGrpcService *scac.Grpc
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) initGRPCServer() *Builder {
	if os.Getenv("PRODUCTION") == "1" {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		b.server = grpc.NewServer(grpc.Creds(tlsCredentials))
	} else {
		b.server = grpc.NewServer()
	}
	return b
}

func (b *Builder) initEnvVariables() *Builder {
	variables, err := getEnvVariables()
	if err != nil {
		panic(err)
		return nil
	}

	b.variables = variables
	return b
}

func (b *Builder) initDatabase() *Builder {
	db, err := getDatabase(b.variables.PostgresConfig)
	if err != nil {
		panic(err)
		return nil
	}
	b.db = db
	return b
}

func (b *Builder) getArgsForTrackers() *tracking.BaseConstructorArgumentsForTracker {
	return &tracking.BaseConstructorArgumentsForTracker{
		Request:            requests.New(),
		UserAgentGenerator: requests.NewUserAgentGenerator(),
		Datetime:           datetime.NewDatetime(),
	}
}

func (b *Builder) initSitcStore() *Builder {
	b.sitcStore = login_provider.NewStore(
		b.variables.SitcServiceUsername,
		b.variables.SitcServicePassword,
		b.variables.SitcServiceBasicAuth,
		b.variables.SitcAccessToken,
	)
	return b
}

func (b *Builder) initCaptchaSolver() *Builder {
	b.captchaSolver = captcha_resolver.NewCaptcha(
		captcha_resolver.NewRandomStringGenerator(),
		captcha_resolver.NewCaptchaGetter(b.getArgsForTrackers().Request, b.getArgsForTrackers().UserAgentGenerator),
		captcha_resolver.NewCaptchaSolver(b.getArgsForTrackers().Request, b.variables.TwoCaptchaApiKey),
	)
	return b
}

func (b *Builder) initUnlocodesRepo() *Builder {
	b.unclocodesRepository = sklu.NewRepository(b.db)
	return b
}

func (b *Builder) initCache() *Builder {
	b.cache = cache.NewCache(redis.NewClient(&redis.Options{
		Addr:     b.variables.RedisConfig.Url,
		Password: "", // no password set
		DB:       0,  // use default DB
	}), parseExpiration(b.variables.RedisConfig.Ttl))
	return b
}

func (b *Builder) initLoginProvider() *Builder {
	b.loginProvider = login_provider.NewTaskManager(
		time.Hour*1,
		login_provider.NewTaskGenerator(b.sitcStore, login_provider.NewProvider(
			b.sitcStore.Username(),
			b.sitcStore.Password(),
			b.sitcStore.BasicAuth(),
			login_provider.NewRequest(b.getArgsForTrackers().Request, b.getArgsForTrackers().UserAgentGenerator),
			b.captchaSolver,
		)),
		scheduler.NewDefault(""),
	)
	return b
}

func (b *Builder) initContainerTrackers() *Builder {
	b.cosuContainerTracker = cosu.NewContainerTracker(b.getArgsForTrackers())
	b.fesoContainerTracker = feso.NewContainerTracker(b.getArgsForTrackers())
	b.haluContainerTracker = halu.NewContainerTracker(b.getArgsForTrackers(), b.unclocodesRepository)
	b.maeuContainerTracker = maeu.NewContainerTracker(b.getArgsForTrackers())
	b.mscuContainerTracker = mscu.NewContainerTracker(b.getArgsForTrackers())
	b.oneyContainerTracker = oney.NewContainerTracker(b.getArgsForTrackers())
	b.sitcContainerTracker = sitc.NewContainerTracker(b.getArgsForTrackers(), b.sitcStore)
	b.skluContainerTracker = sklu.NewContainerTracker(b.getArgsForTrackers(), b.unclocodesRepository)
	return b
}

func (b *Builder) initScacRepository() *Builder {
	b.scacRepository = scac_accessory.NewRepository(b.db)
	return b
}

func (b *Builder) initMainContainerTracker() *Builder {
	b.containerTrackingMainTracker = tracking.NewContainerTracker(map[string]tracking.IContainerTracker{
		"COSU": b.cosuContainerTracker,
		"FESO": b.fesoContainerTracker,
		"HALU": b.haluContainerTracker,
		"MAEU": b.maeuContainerTracker,
		"MSCU": b.mscuContainerTracker,
		"ONEY": b.oneyContainerTracker,
		"SITC": b.sitcContainerTracker,
		"SKLU": b.skluContainerTracker,
	})
	return b
}

func (b *Builder) initContainerTrackingService() *Builder {
	b.containerTrackingService = tracking.NewContainerTrackingService(
		b.containerTrackingMainTracker,
		b.scacRepository,
		logging.NewLogger("containerTrackingService"),
		b.cache,
	)
	return b
}

func (b *Builder) initContainerTrackingGRPCService() *Builder {
	b.containerTrackingGrpcService = tracking.NewContainerTrackingGrpc(
		b.containerTrackingService,
		logging.NewLogger("containerTrackingGrpc"),
	)
	return b
}

func (b *Builder) registerContainerTrackingGRPCService() *Builder {
	pb.RegisterTrackingByContainerNumberServer(b.server, b.containerTrackingGrpcService)
	return b
}

func (b *Builder) initBillTrackers() *Builder {
	b.fesoBillTracker = feso.NewBillTracker(b.getArgsForTrackers())
	b.haluBillTracker = halu.NewBillTracker(b.getArgsForTrackers())
	b.sitcBillTracker = sitc.NewBillTracker(
		b.getArgsForTrackers(),
		sitc.NewBillTrackingRequest(
			b.getArgsForTrackers().Request,
			b.getArgsForTrackers().UserAgentGenerator,
			b.sitcStore,
		),
		b.captchaSolver,
	)
	b.skluBillTracker = sklu.NewBillTracker(b.getArgsForTrackers())
	b.zhguBillTracker = zhgu.NewBillTracker(b.getArgsForTrackers())
	b.reelBillTracking = reel.NewBillTracker(b.getArgsForTrackers())
	return b
}

func (b *Builder) initBillMainTracker() *Builder {
	b.billMainTracker = tracking.NewBillNumberTracker(map[string]tracking.IBillTracker{
		"FESO": b.fesoBillTracker,
		"HALU": b.haluBillTracker,
		"SITC": b.sitcBillTracker,
		"SKLU": b.skluBillTracker,
		"ZHGU": b.zhguBillTracker,
		"REEL": b.reelBillTracking,
	})
	return b
}

func (b *Builder) initBillTrackingService() *Builder {
	b.billTrackingService = tracking.NewBillTrackingService(
		b.billMainTracker,
		b.scacRepository,
		logging.NewLogger("billTrackingService"),
		b.cache,
	)
	return b
}

func (b *Builder) initBillTrackingGRPCService() *Builder {
	b.billTrackingGrpc = tracking.NewBillNumberTrackingGrpc(
		b.billTrackingService,
		logging.NewLogger("billTrackingGrpc"),
	)
	return b
}

func (b *Builder) registerBillTrackingGRPCService() *Builder {
	pb.RegisterTrackingByBillNumberServer(b.server, b.billTrackingGrpc)
	return b
}

func (b *Builder) initScacService() *Builder {
	f, _ := os.Getwd()
	path := fmt.Sprintf(`%s/%s`, filepath.Dir(f), filepath.Base(f))

	file, err := os.ReadFile(fmt.Sprintf(`%s/scac.json`, path))
	if err != nil {
		panic(err)
		return nil
	}
	var s struct {
		ContainerScac []*scac.WithFullName `json:"containerScac"`
		BillScac      []*scac.WithFullName `json:"billScac"`
	}
	if err := json.Unmarshal(file, &s); err != nil {
		panic(err)
		return nil
	}
	b.scacService = scac.NewService(s.ContainerScac, s.BillScac)
	return b
}

func (b *Builder) initScacGrpcService() *Builder {
	b.scacGrpcService = scac.NewGrpc(b.scacService)
	return b
}

func (b *Builder) registerScacGrpcService() *Builder {
	pb.RegisterScacServiceServer(b.server, b.scacGrpcService)
	return b
}

func (b *Builder) Run() {
	if err := b.loginProvider.Run(); err != nil {
		panic(err)
		return
	}
	go func() {
		l, err := net.Listen("tcp", fmt.Sprintf(`0.0.0.0:51372`))
		if err != nil {
			panic(err)
			return
		}
		log.Println("START GRPC SERVER")
		if err := b.server.Serve(l); err != nil {
			panic(err)
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-s
}
