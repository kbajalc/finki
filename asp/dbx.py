import os
import sys
import warnings
warnings.filterwarnings("ignore")

# cdir = os.getcwd()
# if not cdir.endswith("/python"):
#     ndir = os.path.dirname(os.path.abspath('__file__'))
#     pdir = os.path.abspath(os.path.join(ndir, '../../slides'))
#     os.chdir(pdir)
#     # sys.path.insert(0, pdir)
# pass #if

import numpy as np

from pxg import plot
from pxg import Record, MV

# ./exg -dir cat -qrx -bfx -recs -devext -db mitdb

def PlotPage(rec: Record, page = 0, dif = False, simple = False):
    PON, POF = plot.Range(rec, page)

    c = 0
    for i in range(len(rec.Lines)):
        b = rec.Lines[i]
        if b.OTime < PON + 10: continue
        if b.OTime > POF - 10: break

        if ((b.Episode != "AFIB" and b.Class != b.Dlass) or (b.Episode == "AFIB" and b.Class == "V" and b.Dlass != "V")) and b.Class != 'Q': 
            c += 1
            if dif: b.Marker = "DIF"
        # else:
        #     b.Ignore = True
        pass #if
    pass #for
    if dif and c == 0: return

    return plot.Page(
        rec, 
        page, 
        0,

        SENSOR = True,
        SIGNAL = True,
        DECTOR = False,
        EXTEND = "FORTER",

        label = "BPM",
        angle = 0,

        punts = True, 
        rrqs  = True, 
        qsvl  = False, 
        trig  = False, 
        anref = False,
        onoff = False,

        simple = simple,
        marker = dif,
    )
pass #if

def Plot(db: str, rid: str, pages = [], dif = False, simple = False):
    rec = Record(db, rid, dir = "cat")
    if pages == []: pages = range(0, len(rec.Signal) // plot.CHUNK)
    for pg in pages:
        PlotPage(rec, pg, dif, simple)
    pass #for
pass #def
