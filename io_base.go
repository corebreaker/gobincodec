package bincodec

type IoBase interface {
    ToWriter(io.Writer) Writer
    ToWriterAt(io.WriterAt) WriterAt

    ToReader(io.Reader) Reader
    ToReaderAt(io.ReaderAt) ReaderAt
}

type tIoBase struct {
    tCodecBase
}

func (self *tIoBase) ToWriter(w io.Writer) Writer {
    return &tWriter{
        tIoBase: *self,
        w:       w,
    }
}

func (self *tIoBase) ToWriterAt(w io.WriterAt) WriterAt {
    return &tWriterAt{
        tIoBase: *self,
        w:       w,
    }
}

func (self *tIoBase) ToReader(r io.Reader) Reader {
    return &tReader{
        tIoBase: *self,
        r:       r,
    }
}

func (self *tIoBase) ToReaderAt(r io.ReaderAt) ReaderAt {
    return &tReaderAt{
        tIoBase: *self,
        r:       r,
    }
}

func MakeIo() IoBase {
    return &tIoBase{
        tCodecBase: tCodecBase{
            spec: new_spec(),
        },
    }
}
