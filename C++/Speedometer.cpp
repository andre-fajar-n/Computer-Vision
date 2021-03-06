#include "opencv2/imgproc.hpp"
#include "opencv2/imgcodecs.hpp"
#include "opencv2/highgui.hpp"
#include "iostream"

using namespace cv;
using namespace std;

int main(int argc, char** argv)
{
	namedWindow("original", WINDOW_AUTOSIZE);
	namedWindow("result0", WINDOW_AUTOSIZE);
	
	VideoCapture cap; cap.open(0);
	Mat img0; Mat tmp; Mat hasil;
	Rect rec(100, 100, 200, 200);
	double min, max;
	Point minL, maxL;

	for (;;)
	{
		cap >> img0;
		if (img0.empty()) break;
		if (tmp.empty()) tmp = img0(rec).clone();
		matchTemplate(img0, tmp, hasil, TM_SQDIFF);
		minMaxLoc(hasil, &min, &max, &minL, &maxL);
		tmp = img0(rec).clone();
		rectangle(img0, rec, Scalar(0, 0, 255), 3, 8, 0);
		//Mat tmp = img0(rec);
		Mat result = Mat::zeros(Size(img0.cols - tmp.cols + 1, img0.rows - tmp.rows + 1), 1);
		rectangle(img0, Rect(minL.x, minL.y, tmp.cols, tmp.rows), Scalar(0, 255, 0), 2, 8);
		
		int x = minL.x - 100;
		int y = minL.y - 100;
		//minMaxLoc(result, &min, &max, &minL, &maxL);
		//printf("nilai minimum:%f - nilai maksimum:%f\n", min, max);
		printf("lokasi awal:(100, 100) - lokasi akhir:(%d, %d)	-	", minL.x, minL.y);
		printf("kecepatan: (%d,%d)\n", x, y);
		//rectangle(img0, Rect(minL.x, minL.y, tmp.cols, tmp.rows), Scalar(255, 0, 0), 1, 8);
		//rectangle(img1, Rect(minL.x, minL.y, tmp.cols, tmp.rows), Scalar(0, 255, 0), 1, 8);

		imshow("original", img0);
		imshow("result0", tmp);
		
		char pencet = waitKey(10);
		if (pencet == 32) {
			while (pencet == 32) {
				if (waitKey(10) == 32) break; //tombol pause
			}
		}
		else if (pencet == 27) return 0; //tombol exit
	}
}