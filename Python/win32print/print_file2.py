import os, sys
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

# print_data(data) isspausdina teksta
def print_data(data):
	hPrinter = win32print.OpenPrinter(PRINTER)
	try:
	  hJob = win32print.StartDocPrinter (hPrinter, 1, ("pitono skriptas", None, "RAW"))
	  try:
		win32print.StartPagePrinter (hPrinter)
		win32print.WritePrinter (hPrinter, data)
		win32print.EndPagePrinter (hPrinter)
	  except Exception as e:
		print "Exception:", e
	  finally:
		win32print.EndDocPrinter (hPrinter)
	except Exception as e:
		print "Exception:", e
	finally:
	  win32print.ClosePrinter (hPrinter)

# print_file(_filename) atidaro ir isspausdina faila
# t.y. nuskaito faila ir perduoda teksta i print_data(data)
def print_file(_filename):
	with open(_filename, "r") as f:
		data = f.read()
	print_data(data)
	
print_file("myfile.txt")