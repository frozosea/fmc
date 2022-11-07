package excel_writer

import (
	"fmt"
	file_reader "github.com/frozosea/file-reader"
	"github.com/xuri/excelize/v2"
	"schedule-tracking/pkg/tracking"
	"time"
)

type IWriter interface {
	WriteContainerNo(result tracking.ContainerNumberResponse, timeFormatter func(time.Time) string) (string, error)
	WriteBillNo(result tracking.BillNumberResponse, timeFormatter func(time.Time) string) (string, error)
}

type baseWriter struct {
}

func newBaseWriter() *baseWriter {
	return &baseWriter{}
}

func (w *baseWriter) writeUpColoumns(file *excelize.File) error {
	if err := file.SetCellStr("Sheet1", "A1", "Time"); err != nil {
		return err
	}
	if err := file.SetCellStr("Sheet1", "B1", "Operation Name"); err != nil {
		return err
	}
	if err := file.SetCellStr("Sheet1", "C1", "Location"); err != nil {
		return err
	}
	if err := file.SetCellStr("Sheet1", "D1", "Vessel"); err != nil {
		return err
	}
	return nil
}
func (w *baseWriter) writeInfoAboutMoving(file *excelize.File, infoAboutMoving []tracking.BaseInfoAboutMoving, timeFormatter func(time.Time) string) error {
	var index = 2
	for _, value := range infoAboutMoving {
		if err := file.SetCellStr("Sheet1", fmt.Sprintf(`A%d`, index), timeFormatter(value.Time)); err != nil {
			return err
		}
		if err := file.SetCellStr("Sheet1", fmt.Sprintf(`B%d`, index), value.OperationName); err != nil {
			return err
		}
		if err := file.SetCellStr("Sheet1", fmt.Sprintf(`C%d`, index), value.Location); err != nil {
			return err
		}
		if err := file.SetCellStr("Sheet1", fmt.Sprintf(`D%d`, index), value.Vessel); err != nil {
			return err
		}
		index++
	}
	return nil
}
func (w *baseWriter) writeUpColoumnsAndInfoAboutMoving(file *excelize.File, infoAboutMoving []tracking.BaseInfoAboutMoving, timeFormatter func(time.Time) string) error {
	if writeUpColoumnsErr := w.writeUpColoumns(file); writeUpColoumnsErr != nil {
		return writeUpColoumnsErr
	}
	if writeInfoAboutMovingErr := w.writeInfoAboutMoving(file, infoAboutMoving, timeFormatter); writeInfoAboutMovingErr != nil {
		return writeInfoAboutMovingErr
	}
	return nil
}

type Writer struct {
	dirName string
	*baseWriter
	reader *file_reader.FileReader
}

func NewWriter(dirName string) *Writer {
	return &Writer{dirName: dirName, baseWriter: newBaseWriter(), reader: file_reader.New()}
}

func (w *Writer) WriteContainerNo(result tracking.ContainerNumberResponse, timeFormatter func(time.Time) string) (string, error) {
	file := excelize.NewFile()
	if baseWriteErr := w.writeUpColoumnsAndInfoAboutMoving(file, result.InfoAboutMoving, timeFormatter); baseWriteErr != nil {
		return "", baseWriteErr
	}
	filePath := w.reader.GetFileNameByDirNameAndFilename(w.dirName, result.Container)
	if saveErr := file.SaveAs(fmt.Sprintf(`%s.xlsx`, filePath)); saveErr != nil {
		return "", saveErr
	}
	return fmt.Sprintf(`%s.xlsx`, filePath), nil
}

func (w *Writer) WriteBillNo(result tracking.BillNumberResponse, timeFormatter func(time.Time) string) (string, error) {
	file := excelize.NewFile()
	if baseWriteErr := w.writeUpColoumnsAndInfoAboutMoving(file, result.InfoAboutMoving, timeFormatter); baseWriteErr != nil {
		return "", baseWriteErr
	}
	if err := file.SetCellStr("Sheet1", fmt.Sprintf(`A%d`, len(result.InfoAboutMoving)+3), "ETA"); err != nil {
		return "", err
	}
	if err := file.SetCellStr("Sheet1", fmt.Sprintf(`B%d`, len(result.InfoAboutMoving)+3), timeFormatter(result.EtaFinalDelivery)); err != nil {
		return "", err
	}
	filePath := w.reader.GetFileNameByDirNameAndFilename(w.dirName, result.BillNo)
	if saveErr := file.SaveAs(fmt.Sprintf(`%s.xlsx`, filePath)); saveErr != nil {
		return "", saveErr
	}
	return fmt.Sprintf(`%s.xlsx`, filePath), nil
}
