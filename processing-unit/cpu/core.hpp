#ifndef go_opencv_cpu_core
#define go_opencv_cpu_core

#include "../../entity/entity.hpp"

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif // __cplusplus

#ifdef __cplusplus
typedef cv::Mat* Mat;
#else
typedef void* Mat;
#endif // __cplusplus

Mat NewMat();
Mat NewMatFromBytes(int rows, int cols, int type, struct ByteArray buf);

#ifdef __cplusplus
}
#endif // __cplusplus

#endif //go_opencv_cpu_core