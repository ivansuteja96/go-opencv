
#include "entity.hpp"

bool IsCudaAvailable() {
    int num_devices = cv::cuda::getCudaEnabledDeviceCount();
    if(num_devices == 0) {
        return false;
    } else {
        return true;
    }
}
