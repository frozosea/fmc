package main

import (
	"errors"
	"fmt"
	"os"
)

type EnvVariables struct {
	UserAppIp                     string
	UserAppPort                   string
	AltsKeyForUserApp             string
	TrackingAppIp                 string
	TrackingAppPort               string
	AltsKeyForTrackingApp         string
	ScheduleTrackingIp            string
	ScheduleTrackingPort          string
	AltsKeyForScheduleTrackingApp string
}

func getEnvVariable(variableName string) (string, error) {
	variable := os.Getenv(variableName)
	if variable == "" {
		return "", errors.New(fmt.Sprintf(`no %s env variable`, variableName))
	}
	return variable, nil
}

func getEnvVariables() (*EnvVariables, error) {
	variables := map[string]string{
		"USER_APP_IP":            "",
		"USER_APP_PORT":          "",
		"TRACKING_IP":            "",
		"TRACKING_PORT":          "",
		"SCHEDULE_TRACKING_HOST": "",
		"SCHEDULE_TRACKING_PORT": "",
		"FREIGHT_HOST":           "",
		"FREIGHT_PORT":           "",
	}
	for name := range variables {
		v, err := getEnvVariable(name)
		if err != nil {
			return nil, err
		}
		variables[name] = v
	}
	return &EnvVariables{
		UserAppIp:            variables["USER_APP_IP"],
		UserAppPort:          variables["USER_APP_PORT"],
		TrackingAppIp:        variables["TRACKING_IP"],
		TrackingAppPort:      variables["TRACKING_PORT"],
		ScheduleTrackingIp:   variables["SCHEDULE_TRACKING_HOST"],
		ScheduleTrackingPort: variables["SCHEDULE_TRACKING_PORT"],
	}, nil
}

func getUrl(ip, port string) string {
	return fmt.Sprintf("%s:%s", ip, port)
}
