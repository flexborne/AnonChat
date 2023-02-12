#pragma once

#include <QObject>

#include "ServerCommand.h"

#include "enum_AgeCategory.h"
#include "enum_ConversationCategory.h"
#include "enum_Sex.h"
#include "utils_bitwiseEnum.h"


struct User : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString id READ getId NOTIFY idChanged)

public:
    Sex::e sex;
    AgeCategory::e ageCategory;
    ConversationCategory::e conversationCategory;

    struct Preferences
    {
        Sex::e sex;
        Flags<AgeCategory::e> ageFlags;
    } preferences;

    QString getId() const noexcept;

signals:
    void findMatch(CmdFindMatch);
    void idChanged();

public slots:
    void set(Sex::e sex, AgeCategory::e ageCategory, ConversationCategory::e conversationCategory,
             Sex::e sexPreferences, AgeCategory::e agePreferences);

    void setId(std::string id);

private:
    QString id;
};

