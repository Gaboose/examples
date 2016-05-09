import win32print

PRINTER = win32print.GetDefaultPrinter()

print
print "Name:", PRINTER
print "Escaped Name: \"%s\"" % PRINTER.replace("\\", "\\\\")