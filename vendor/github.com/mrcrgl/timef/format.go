package timef

import (
	"errors"
	"math"
	"time"
)

const (
	digits         = "0123456789"
	octHyphen      = byte('-')
	octLetterT     = byte('T')
	octColon       = byte(':')
	octLetterPlus  = byte('+')
	octLetterMinus = byte('-')
)

const (
	RFC3339BufSize = 25
	RFC3339BufCap  = 25
)

var ErrNotSupported = errors.New("FormatNotSupported")

func Format(layout string, t time.Time) (string, error) {
	s, err := FormatBytes(layout, t)
	return string(s), err
}

func FormatBytes(layout string, t time.Time) ([]byte, error) {
	switch layout {
	case time.RFC3339:
		return FormatRFC3339(t), nil
	}

	return nil, ErrNotSupported
}

// "2006-01-02T15:04:05Z07:00"
func FormatRFC3339(t time.Time) []byte {
	bs := make([]byte, RFC3339BufSize, RFC3339BufCap)
	WriteRFC3339(t, bs)

	return bs
}

func WriteRFC3339(t time.Time, bs []byte) (int, error) {
	return WriteRFC3339At(t, bs, 0)
}

func WriteRFC3339At(t time.Time, bs []byte, off int64) (int, error) {
	b := &buffer{t: bs, p: off}

	year, month, day := t.Date()
	hour, minute, second := t.Clock()

	_, offset := t.Zone()
	var sign byte
	if math.Signbit(float64(offset)) {
		sign = octLetterMinus
		offset *= -1
	} else {
		sign = octLetterPlus
	}

	offsetMin := offset / 60

	b.writeNDigits(4, year, ' ')
	b.writeByte(octHyphen)
	b.writeTwoDigits(int(month))
	b.writeByte(octHyphen)
	b.writeTwoDigits(day)
	b.writeByte(octLetterT)
	b.writeTwoDigits(hour)
	b.writeByte(octColon)
	b.writeTwoDigits(minute)
	b.writeByte(octColon)
	b.writeTwoDigits(second)
	b.writeByte(sign)
	b.writeTwoDigits(offsetMin / 60)
	b.writeByte(octColon)
	b.writeTwoDigits(offsetMin % 60)

	return int(b.p - off), nil
}

type buffer struct {
	t []byte
	p int64
}

func (b *buffer) writeByte(a ...byte) {
	copy(b.t[b.p:b.p+int64(len(a))], a)
	b.p += int64(len(a))
}

func (b *buffer) writeTwoDigits(i int) {
	b.t[b.p+1] = digits[i%10]
	i /= 10
	b.t[b.p] = digits[i%10]
	b.p += 2
}

func (b *buffer) writeNDigits(n int64, d int, pad byte) {
	j := n - 1
	for ; j >= 0 && d > 0; j-- {
		b.t[b.p+j] = digits[d%10]
		d /= 10
	}
	for ; j >= 0; j-- {
		b.t[b.p+j] = pad
	}
	b.p += n
}
