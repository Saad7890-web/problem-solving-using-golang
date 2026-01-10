package leetcodecontest
func largestEven(s string) string {
    lastIndex := -1

    for i := 0; i < len(s); i++{
        if s[i] == '2'{
            lastIndex = i
        }
    }
    if lastIndex == -1 {
        return ""
    }
    return s[:lastIndex+1]
}