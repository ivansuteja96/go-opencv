#include <stdbool.h>

#ifndef go_opencv_entity
#define go_opencv_entity

typedef struct ByteArray {
    char* data;
    int length;
} ByteArray;

#ifdef __cplusplus
#include <opencv2/core/cuda.hpp>
extern "C" {
#endif // __cplusplus

bool IsCudaAvailable();

#ifdef __cplusplus
}
#endif // __cplusplus

#endif // go_opencv_entity
