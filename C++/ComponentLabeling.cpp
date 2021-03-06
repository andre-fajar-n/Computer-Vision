#include "opencv/cv.h"
#include "opencv/highgui.h"

int findcontour(IplImage* img)
{
	int koin = 0;
	int uang[4];
	CvMemStorage* storage = cvCreateMemStorage();
	CvSeq*first_contour = NULL;
	cvClearMemStorage(storage);
	int Nc = cvFindContours(img, storage, &first_contour, sizeof(CvContour), CV_RETR_LIST);
	int count = 0;
	for (CvSeq* c = first_contour; c != NULL; c = c->h_next)
	{
		if (c->total > 40)
		{
			count++;
			printf("%d-", c->total);
			if (c->total > 120 && c->total < 150) {
				printf("koin 500 : %d \n", c->total);
				koin += 500;
			}
			else if (c->total > 95 && c->total <= 120) {
				printf("koin 200 : %d \n", c->total);
				koin += 200;
			}
			else if (c->total > 80 && c->total <= 95) {
				printf("koin 100 : %d \n", c->total);
				koin += 100;
			}
			/*for (int i = 0; i < c->total; ++i)
			{
			 CvPoint* p = CV_GET_SEQ_ELEM(CvPoint, c, i);
			 p->x;
			 p->y;
			}*/
		}
	}

	cvClearMemStorage(storage);
	return count;
}

int main()
{
	IplImage* gambar = 0;
	IplImage* gambarnow;
	CvCapture* capture = cvCaptureFromCAM(1);
	cvNamedWindow("windowx", 1);
	for (;;)
	{
		//Ambil GambarNow
		gambar = cvQueryFrame(capture);
		if (gambar == 0)break;
		gambarnow = cvCreateImage(cvSize(gambar->width, gambar->height), 8, 1);
		cvCvtColor(gambar, gambarnow, CV_BGR2GRAY);
		cvThreshold(gambarnow, gambarnow, 70, 255, CV_THRESH_BINARY_INV);
		cvShowImage("windowx", gambarnow);
		if (cvWaitKey(1) > 0)break;
		//Program Connected Component Labelling
		int jumlah = findcontour(gambarnow);

		printf("Jumlah Blob  %d\n", jumlah);
	}
	cvReleaseImage(&gambar);
	cvReleaseImage(&gambarnow);
	cvReleaseCapture(&capture);
	return 0;
}