function sexToString(enumValue) {
    switch(enumValue) {
        case Sex.Male:
            return "Male";
        case Sex.Female:
            return "Female";
        case Sex.Undefined:
            return "Undefined";
        default:
            return "";
    }
}

function ageCatToString(enumValue) {
    switch(enumValue) {
        case AgeCategory.Under18:
            return "Under 18";
        case AgeCategory.Between18And25:
            return "Between 18 and 25";
        case AgeCategory.Between26And35:
            return "Between 26 and 35";
        case AgeCategory.Over35:
            return "Over 35";
        default:
            return "";
    }
}


function conversationCatToString(enumValue) {
    switch(enumValue) {
        case ConversationCategory.Casual:
            return "Casual";
        case ConversationCategory.Serious:
            return "Serious";
        case ConversationCategory.Flirtatious:
            return "Flirtatious";
        default:
            return "";
    }
}
