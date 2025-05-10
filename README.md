# TENX Technical test

## Create function to classified if the image is colorful or not
the logic is to identify pixels by pixels check if the pixels provided is coloful or not, need to check it one by one
what to do is 
for example R G B (122, 125, 0)
### 1. need to know min and max value from the pixels
`` 
(max value of the pixels -  min value of the pixels) /  (max value of the pixels 
(125 - 0) / 125 = 1
``
### 2. define threshold to identify if the pixel is classified as colorful or not
if the threshold 0.2, 
so we just need to compare it 1 >= 0.2 (?) 
if value from the pixels is higher than the threshold it tells that the pixels is colorful

### 3. define image is colorful or not
count how much pixels is colorful, for example if there is 25 pixels, and 17 of it is colorful so we just need to 25 / 17 = 0.68, again, define the threshold, 
if the threshold is 0.4 so we just need to compare it
0.68 > 0.4 (?) 
if its bigger then the image is classified as colorful, so is the grayscale function it same but different in some part, just reverse
