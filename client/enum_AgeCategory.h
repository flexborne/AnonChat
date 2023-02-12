#pragma once

#include <QObject>

namespace AgeCategory {
    Q_NAMESPACE
    enum class e {
        Under18 = 1 << 0,
        Between18And25 = 1 << 1,
        Between26And35 = 1 << 2,
        Over35 = 1 << 3
    };
    Q_ENUM_NS(e)
}
