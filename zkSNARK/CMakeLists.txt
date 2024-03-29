#cmake_minimum_required(VERSION 3.9)
#project(zkSNARK)
#
#set(CMAKE_CXX_STANDARD 11)
#
#add_executable(zkSNARK main.cpp)
cmake_minimum_required(VERSION 2.8)

project(zkSNARK)

include_directories(/usr/local/opt/openssl/include)
#target_link_libraries (libsnark /usr/local/opt/openssl/lib)
LINK_DIRECTORIES(/usr/local/opt/openssl/lib)
set(ENV{PKG_CONFIG_PATH} "/usr/local/opt/openssl/lib/pkgconfig")

set(
        CURVE
        "ALT_BN128"
        CACHE
        STRING
        "Default curve: one of ALT_BN128, BN128, EDWARDS, MNT4, MNT6"
)

set(
        OPT_FLAGS
        ""
        CACHE
        STRING
        "Override C++ compiler optimization flags"
)

option(
        MULTICORE
        "Enable parallelized execution, using OpenMP"
        OFF
)

option(
        WITH_PROCPS
        "Use procps for memory profiling"
        OFF
)

option(
        WITH_SUPERCOP
        "Support for Ed25519 signatures required by ADSNARK"
        OFF
)

option(
        VERBOSE
        "Print internal messages"
        OFF
)

if(CMAKE_COMPILER_IS_GNUCXX OR "${CMAKE_CXX_COMPILER_ID}" STREQUAL "Clang")
    # Common compilation flags and warning configuration
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -Wall -Wextra -Wfatal-errors -pthread")

    if("${MULTICORE}")
        set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fopenmp")
    endif()

    # Default optimizations flags (to override, use -DOPT_FLAGS=...)
    if("${OPT_FLAGS}" STREQUAL "")
        set(OPT_FLAGS "-ggdb3 -O2 -march=native -mtune=native")
    endif()
endif()

add_definitions(-DCURVE_${CURVE})

if(${CURVE} STREQUAL "BN128")
    add_definitions(-DBN_SUPPORT_SNARK=1)
endif()

if("${VERBOSE}")
    add_definitions(-DVERBOSE=1)
endif()

if("${MULTICORE}")
    add_definitions(-DMULTICORE=1)
endif()

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} ${OPT_FLAGS}")

include(FindPkgConfig)
if("${WITH_PROCPS}")
    pkg_check_modules(PROCPS REQUIRED libprocps)
else()
    add_definitions(-DNO_PROCPS)
endif()

include_directories(.)

add_subdirectory(depends)
add_subdirectory(src)