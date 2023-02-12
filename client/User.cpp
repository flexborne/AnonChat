#include "User.h"

#include <iostream>


QString User::getId() const noexcept
{
    return id;
}

void User::set(Sex::e sex, AgeCategory::e ageCategory, ConversationCategory::e conversationCategory, Sex::e sexPreferences, AgeCategory::e agePreferences)
{
    this->sex = sex;
    this->ageCategory = ageCategory;
    this->conversationCategory = conversationCategory;
    preferences.sex = sexPreferences;
    preferences.ageFlags = agePreferences;

    emit findMatch(CmdFindMatch{*this});
}

void User::setId(std::string id)
{
    this->id = QString::fromStdString(std::move(id));
}
