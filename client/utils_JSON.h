#pragma once

#include <Poco/JSON/Object.h>
#include <Poco/JSON/Stringifier.h>
#include <Poco/JSON/Parser.h>

#include <optional>

#include <string>
#include <QDebug>

template <class T, class U>
void createJSONHelper(Poco::JSON::Object& json, const U& tag, const T& value)
{
    json.set(tag, value);
}

template <class T, class U, class ... Args>
void createJSONHelper(Poco::JSON::Object& json, const U& tag, const T& value, const Args& ... args)
{
    createJSONHelper(json, tag, value);
    createJSONHelper(json, args ...);
}

template <class ... Args>
auto createJSON(const Args& ... args) -> std::string
{
    Poco::JSON::Object json;
    createJSONHelper(json, args ...);

    std::ostringstream stream;
    Poco::JSON::Stringifier::stringify(json, stream);

    return stream.str();
}

template <class T>
auto getValueFromJSON(const std::string& json, const auto& tag) -> std::optional<T> try
{
    Poco::JSON::Parser parser;

    // parse the JSON string
    Poco::Dynamic::Var result = parser.parse(json);

    // access the values
    Poco::JSON::Object::Ptr object = result.extract<Poco::JSON::Object::Ptr>();

    return object->getValue<T>(std::string{tag});
} catch (Poco::Exception exception) {
    qDebug() << exception.what();
    return std::nullopt;
}

auto createSimpleChatMessageJSON(const std::string& message) -> std::string;
