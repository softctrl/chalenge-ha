# -*- coding: utf-8 -*-

"""
scgen.py
I am developing this program to help-me at my work.
I hope that this project may help you too.
But, i will not be responsible for any trouble that you may have.
You are by your own.
If you do not have knowledge so is better you do not use this.
Copyright (c) 2017 Carlos Timoshenko Rodrigues Lopes carlostimoshenkorodrigueslopes@gmail.com
http://www.0x09.com.br
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
"""
if __name__ == '__main__':
    MAX = 50000000  # Change for whatever you like
    GROW_RATE = 0.1   # Grow rate
    LATITUDE = 51.925146 + GROW_RATE
    LONGITUDE = 4.478617 + GROW_RATE
    print('"id","lat","lng"')
    for id in range(0, MAX):
        print("%d,%f,%f" % (id + 1, LATITUDE, LONGITUDE))
        LATITUDE = LATITUDE + GROW_RATE
        LONGITUDE = LONGITUDE + GROW_RATE
