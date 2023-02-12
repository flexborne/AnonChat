#pragma once

#include <type_traits>

template <typename E>
constexpr typename std::underlying_type<E>::type to_underlying(E e) noexcept {
    return static_cast<typename std::underlying_type<E>::type>(e);
}

template<typename Enum>
class Flags
{
public:
    using underlying_type = typename std::underlying_type<Enum>::type;

    constexpr Flags() = default;
    constexpr Flags(Enum flag) : value_(to_underlying(flag)) {}
    constexpr Flags(underlying_type value) : value_(value) {}

    template <typename... Args>
    constexpr Flags(Enum flag, Args... args) : value_(to_underlying(flag) | Flags(args...).value_) {}

    constexpr void setFlag(Flags<Enum> flags)
    {
        value_ |= flags.value_;
    }

    constexpr void clearFlag(Enum flag)
    {
        value_ &= ~to_underlying(flag);
    }

    constexpr bool hasFlag(Enum flag) const
    {
        return (value_ & to_underlying(flag)) != 0;
    }

    constexpr Flags operator|(Enum rhs) const
    {
        return Flags(value_ | to_underlying(rhs));
    }

    constexpr Flags operator&(Enum rhs) const
    {
        return Flags(value_ & to_underlying(rhs));
    }

    constexpr Flags operator^(Enum rhs) const
    {
        return Flags(value_ ^ to_underlying(rhs));
    }

    constexpr Flags operator~() const
    {
        return Flags(~value_);
    }

    constexpr Flags& operator|=(Enum rhs)
    {
        value_ |= to_underlying(rhs);
        return *this;
    }

    constexpr Flags& operator&=(Enum rhs)
    {
        value_ &= to_underlying(rhs);
        return *this;
    }

    constexpr Flags& operator^=(Enum rhs)
    {
        value_ ^= to_underlying(rhs);
        return *this;
    }

    constexpr operator underlying_type() const
    {
        return value_;
    }

    friend bool operator&(const Flags& lhs, Enum rhs) { return (lhs.value_ & static_cast<underlying_type>(rhs)) != 0; }
    friend Flags operator&(Enum lhs, const Flags& rhs) { return Flags(static_cast<underlying_type>(lhs) & rhs.value_); }
    friend Flags operator|(Enum lhs, const Flags& rhs) { return Flags(static_cast<underlying_type>(lhs) | rhs.value_); }
    friend Flags operator|(const Flags& lhs, Enum rhs) { return Flags(lhs.value_ | static_cast<underlying_type>(rhs)); }

private:
    underlying_type value_ = 0;
};
