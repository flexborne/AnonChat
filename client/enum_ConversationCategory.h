#pragma once

#include <QObject>

namespace ConversationCategory {
    Q_NAMESPACE
    enum class e {
        Casual,
        Serious,
        Flirtatious,
    };
    Q_ENUM_NS(e)
}
