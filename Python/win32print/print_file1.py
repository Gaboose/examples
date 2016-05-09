import win32api
import win32print

# Kuri printeri naudot:
PRINTER = win32print.GetDefaultPrinter()

# Paleisk "python default_printer.py" kad suzinotum kuris yra defaultinis
#
# Jei default_printer.py isspausdina ne ta printeri, kuri nori naudot,
# pakeisk 'printer' kintamaji i 'Escaped Name' laukeli kuri rodo
# "python list_printers.py"
#
# Pvz as rasyciau:
# PRINTER = "\\\\SC-STUDSPOOLPS\\PULLPRINT"
#
# Nes ant mano kompo "python list_printers.py connections" isspausdina sitai:
#
# Name: \\SC-STUDSPOOLPS\PULLPRINT
# Escaped Name: "\\\\SC-STUDSPOOLPS\\PULLPRINT"
# Description: \\SC-STUDSPOOLPS\PULLPRINT,Canon Generic PCL6 Driver,
# Comment:
#

def print_file(_filename):
    win32api.ShellExecute(0,"print",_filename,'/D:"%s"'%PRINTER,".",0)

print_file("myfile.txt")

