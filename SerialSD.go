package deej-display

import (
	"github.com/jax-b/deej"
	"go.uber.org/zap"
	"io/ioutil"
)
type SerialSD struct {
	sio			 *Deej.SerialIO
	logger  *zap.SugaredLogger
}

// NewSerialSD
func NewSerialSD(sio *Deej.SerialIO) (*serSD, error) {
	logger = sio.logger.Named("SD")
	sersd := &SerialSD{
		sio:	sio,
		logger:	logger,
	}
	return serSD, nil
}

// ListDir lists the dir to logger and returns it as a string
func (serSD *SerialSD) ListDir() (string, error) {
	serSD.sio.Pause()
	serSD.logger.info("SDCard File list")
	serSD.sio.WriteStringLine(serSD.logger, "deej.modules.sd.list")
	SerialData := <-serSD.sio.readLine(serSD.logger)
	for SerialData != "" {
		serSD.logger.info(SerialData)
		SerialData = <-serSD.sio.readLine(serSD.logger)
	}
	serSD.sio.Start()
}

// Delete deletes a file off of the SD card
func (serSD *SerialSD) Delete(filename string) (string, error) {
	serSD.sio.Pause()

	serSD.logger.Infof("Deleting %q from the SD Card", filename)
	serSD.sio.WriteStringLine(serSD.logger, "deej.modules.sd.delete")
	serSD.sio.WriteStringLine(serSD.logger, filename)

	success, cmdKey := serSD.sio.WaitFor(serSD.logger "FILEDELETED")

	if success == false{
		serSD.logger.ErrorW("Failed to delete file", filename)
		return cmdKey
	}

	serSD.sio.Start()
}

func (serSD *SerialSD) SendFile(filepath string, DestFilename string) error {
	serSD.sio.Pause()

	serSD.logger.Infof("Sending %q to the SD Card with %q as the file name", filepath, DestFilename)
	serSD.sio.WriteStringLine(serSD.logger, "deej.modules.sd.send")
	serSD.sio.WriteStringLine(serSD.logger, DestFilename)
	
	f, err := os.open(filepath)
	defer f.Close()


	b1 := make([]byte, 1)
	n1, err := f.Read(b1)
	for n1 > 0 {
		serSD.sio.conn.Write(b1)
		n1, err = f.Read(b1)
	}
	serSD.sio.WriteStringLine("EOF")
	
	serSD.sio.Start()
}