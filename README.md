# golodash
A utility library in Go.

## Install
```
go get github.com/schoeu/goladash
```

## Full list of functions;

Utils:
```
func GetCwd() string
func HasValue(str *[]string, r string) (bool, string)
func GetCurrentDate(date time.Time) string
func Lunar(date string) string
```

IS:
```
func Range(value, left, right float64)
func Email(s string) bool
func URL(str string) bool
func RequestURL(rawurl string) bool
func RequestURI(rawurl string) bool
func Alpha(s string) bool
func UTFLetter(str string) bool
func Alphanumeric(s string) bool
func UTFLetterNumeric(s string) bool
func Numeric(s string) bool
func UTFNumeric(s string) bool
func Whole(value float64) bool
func Natural(value float64) bool
func UTFDigit(s string) bool
func Hexadecimal(str string) bool
func Hexcolor(str string) bool
func RGBcolor(str string) bool
func LowerCase(str string) bool
func UpperCase(str string) bool
func Int(str string) bool
func Float(str string) bool
func ByteLength(str string, min, max int) bool
func UUIDv3(str string) bool
func UUIDv4(str string) bool
func UUIDv5(str string) bool
func UUID(str string) bool
func JSON(str string) bool
func Multibyte(s string) bool
func ASCII(s string) bool
func PrintableASCII(s string) bool
func FullWidth(str string) bool
func HalfWidth(str string) bool
func VariableWidth(str string) bool
func Base64(s string) bool
func IP(str string) bool
func Port(str string) bool 
func IPv4(str string) bool
func IPv6(str string) bool
func MAC(str string) bool
func MongoID(str string) bool
func Latitude(str string) bool
func Longitude(str string) bool
func Semver(str string) bool
func StringLength(str string, min int, max int) bool
func Exists(path string) (bool, error)
func toFloat(str string) (float64, error)
func toInt(str string) (int64, error) 
```

## LICENSE
Copyright (c) 2018 Schoeu

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

