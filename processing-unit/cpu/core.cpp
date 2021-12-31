
#include "../../entity/entity.hpp"
#include "core.hpp"
#include <string.h>

// NewMat creates a new empty Mat
Mat NewMat() {
    return new cv::Mat();
}

Mat NewMatFromBytes(int rows, int cols, int type, struct ByteArray buf) {
    return new cv::Mat(rows, cols, type, buf.data);
}