#include "opencv2/imgproc.hpp"
#include "opencv2/imgcodecs.hpp"
#include "opencv2/highgui.hpp"
#include "iostream"

using namespace cv;
using namespace std;

int manusia = 0, motor = 0;

int blob(Mat img)
{
	vector<vector< Point >> contours;
	vector<Vec4i> hierarchy;
	findContours(img, contours, hierarchy, RETR_LIST, CHAIN_APPROX_SIMPLE);
	int blob = contours.size();
	int jml_blob=0,n=0,a=0;
	for (int n = 0; n < blob; n++)
	{
		if (contours[n].size() > 75)
		{
			jml_blob++;
			a = contours[n].size();
			//cout << "blob =" << jml_blob << "    ukuran blob =" << contours[n].size() << "\n" << endl;
		}
		
	}
	return a;
}

int main(int argc, char** argv)
{
	namedWindow("original", WINDOW_AUTOSIZE);
	namedWindow("result", WINDOW_AUTOSIZE);

	VideoCapture cap("../data/vehicleCounter.avi"); //cap.open(0);
	Mat frame; Mat gray; Mat it0; Mat it1; Mat hasil;
	Mat element2 = getStructuringElement(MORPH_RECT, Size(3, 3), Point(-1, -1));
	Mat element3 = getStructuringElement(MORPH_RECT, Size(35, 35), Point(-1, -1));
	int count = 0;
	Scalar rata2;
	for (;;)
	{
		cap >> frame;
		if (frame.empty()) break;
		cvtColor(frame, gray, COLOR_BGR2GRAY);
		
		//step 1
		it0 = gray.clone();
		if (count == 0)
		{
			it1 = it0.clone();
			count = 1;
			hasil = it0.clone();
		}
		
		//step 2
		absdiff(it0, it1, hasil);
		
		//step 3
		it1 = it0.clone();
		threshold(hasil, hasil, 45, 255, THRESH_BINARY);
		erode(hasil, hasil, element2);
		dilate(hasil, hasil, element3);
		Moments m = moments(hasil, true);
		Point p(m.m10 / m.m00, m.m01 / m.m00);
		int x = (int)(p).x;
		int y = (int)(p).y;
		circle(frame, p, 5, Scalar(0, 0, 255), -1);
		int pixel = blob(hasil);

		if (pixel > 75 && pixel < 110) {
			if (x > 295 && x < 310) {
				manusia++;
			}
		}
		if (pixel > 110 && pixel < 150) {
			if (x > 290 && x < 310) {
				motor++;
			}
		}

		printf("manusia:%d - motor:%d\n", manusia, motor);
		imshow("original", frame);
		imshow("result", hasil);
		char pencet = waitKey(10);
		if (pencet == 32) {
			while (pencet == 32) {
				if (waitKey(10) == 32) break; //pause
			}
		}
		else if (pencet == 27) return 0; //exit
	}
}