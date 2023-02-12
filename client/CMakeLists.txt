cmake_minimum_required(VERSION 3.16)

project(client VERSION 0.1 LANGUAGES CXX)

set(CMAKE_AUTOMOC ON)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_FLAGS "-fPIC")

find_package(Qt6 6.2 REQUIRED COMPONENTS Core Gui Quick)

# Declare the resource file
set(qrc_file project_resources.qrc)

qt6_add_resources(qrc_file_out ${qrc_file})

qt6_wrap_cpp(MOC_FILES enum_AgeCategory.h enum_ConversationCategory.h enum_Sex.h)

add_executable(client main.cpp WebSocketHandler.cpp
    ReceiveRunnable.cpp utils_JSON.cpp ServerCommand.cpp ${qrc_file_out}
    User.cpp ${MOC_FILES})


target_link_libraries(client Qt6::Core Qt6::Gui Qt6::Quick)

# Find Poco
find_package(Poco REQUIRED COMPONENTS Foundation Net JSON)
include_directories(${Poco_INCLUDE_DIRS})
target_link_libraries(client ${Poco_LIBRARIES})

# Find OpenSSL
set(OPENSSL_ROOT_DIR "/System/Volumes/Data/opt/homebrew/Cellar/openssl@3/3.0.7")
find_package(OpenSSL REQUIRED)
include_directories(${OPENSSL_INCLUDE_DIR})
target_link_libraries(client ${OPENSSL_LIBRARIES})

# Find Boost
find_package(Boost REQUIRED COMPONENTS fiber)
include_directories(${Boost_INCLUDE_DIRS})
target_link_libraries(client ${Boost_LIBRARIES})
