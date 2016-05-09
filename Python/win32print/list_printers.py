import win32print
import sys
import re

def list_printers(flag):
	printers = win32print.EnumPrinters(flag)
	if len(printers) == 0:
		print "None"
		return
	
	for i, printer in enumerate(printers):
		print
		print "Name:", printer[2]
		print "Escaped Name: \"%s\"" % printer[2].replace("\\", "\\\\")
		print "Description:", printer[1]
		print "Comment:", printer[3]
	
def main():
	if len(sys.argv) == 1:
		print """
Usage: python %s <option>

Options:
	connections
	container
	default
	expand
	icon<1-8>
	local
	name
	network
	remote
	shared

For example:
	python %s connections""" % (sys.argv[0], sys.argv[0])
		return


	opt = sys.argv[1].upper()
	flagstr = "PRINTER_ENUM_" + opt
	flag = getattr(win32print, flagstr, None)
	if flag == None:
		print "Unknown option: " + opt
		return
	
	list_printers(flag)

main()