from displays.pngDisplay import pngDisplay
from layers.text import TextDrawer

display = pngDisplay()

frame = display.getEmptyFrame()

bigFont = TextDrawer('big')

bigFont.drawString('1234567', frame, (255,255,255))

display.renderFrame(frame)