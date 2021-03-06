#include "opencv/cv.h"
#include "opencv/highgui.h"

int main()
{
	IplImage* gambar = 0;
	/*program menampilkan gambar
	gambar = cvLoadImage("img.jpg", CV_LOAD_IMAGE_COLOR);
	if (gambar > 0)
	{
		cvNamedWindow("display", 1);
		cvShowImage("display", gambar);
		if (cvWaitKey(0)>0)break;	//pencet keyboard langsung close
	}
	*/
	CvCapture* capture = cvCaptureFromCAM(0);
	cvShowImage("display", gambar);
	IplImage* grey;
	for (;;)
	{
		gambar = cvQueryFrame(capture);
		if (gambar > 0)
		{
			grey = cvCreateImage(cvGetSize(gambar), 8, 1);
			cvCvtColor(gambar, grey, CV_RGB2GRAY);
			cvThreshold(grey, grey, 200, 255, CV_THRESH_TOZERO_INV);
			cvThreshold(grey, grey, 50, 255, CV_THRESH_BINARY);
			/*cvAdaptiveThreshold(grey, grey, 255, 0, 0, 31, 5.0);*/
			if (cvWaitKey(1) > 0)break;	//pencet keyboard langsung close
			cvShowImage("display", grey);
		}
	}
	cvReleaseImage(&gambar);
	cvReleaseImage(&grey);
	cvDestroyWindow("display");
	return 0;
}