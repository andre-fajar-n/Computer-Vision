#include "opencv4/opencv2/imgproc.hpp"
#include "opencv4/opencv2/imgcodecs.hpp"
#include "opencv4/opencv2/highgui.hpp"
#include "iostream"

using namespace cv;
using namespace std;


int main(int argc, char** argv)
{
	namedWindow("gambar 1", WINDOW_AUTOSIZE);
	namedWindow("gambar 2", WINDOW_AUTOSIZE);
	namedWindow("gabungan", WINDOW_AUTOSIZE);
	Mat hasil;
	double min, max;
	Point minL, maxL, B;
	int ax, ay;
	Mat img1 = imread(("../data/imageRegistration-1.jpg"), IMREAD_COLOR);
	Mat img2 = imread(("../data/imageRegistration-2.jpg"), IMREAD_COLOR);
	int jml_kolom = img1.cols / 4;
	int jml_baris = img1.rows / 4;
	for (int i = 0; i < 4; i++) {
		for (int j = 0; j < 4; j++) {
			Rect rec(i*jml_kolom, j*jml_baris, jml_kolom, jml_baris);
			Mat tmp = img1(rec).clone();
			matchTemplate(img2, tmp, hasil, TM_SQDIFF);
			minMaxLoc(hasil, &min, &max, &minL, &maxL);
			
			if (min < 2.4e+06) {
				rectangle(img1, rec, Scalar(0, 0, 255), 1, 8);
				rectangle(img2, Rect(minL.x, minL.y, jml_kolom, jml_baris), Scalar(0, 0, 255), 1, 8);
				ax = i * jml_kolom;
				ay = j * jml_baris;
				B.x = minL.x;
				B.y = minL.y;
				cout << "min:" << min << ",	max:" << max << "	(" << ax << "," << ay << ")" << endl;
			}
			imshow("gambar 1", img1);
			imshow("gambar 2", img2);
			//waitKey(0);
		}
	}
	//cout << ax << "-" << ay << "-" << img2.cols << "-" << img2.rows << "-" << minL.x << "-" << minL.y << endl;
	Mat gabungan = Mat::zeros(Size(ax + img2.cols - B.x, ay + img2.rows - B.y),CV_8UC3);
	//cout << "\n" << gabungan.cols << "-" << gabungan.rows << endl;
	img1.copyTo(gabungan(Rect(0, 0, img1.cols, img1.rows)));
	img2.copyTo(gabungan(Rect(ax - B.x, ay - B.y, img2.cols, img2.rows)));
	imshow("gabungan", gabungan);
	waitKey(0);
}