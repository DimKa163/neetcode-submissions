func isAnagram(s string, t string) bool {
 if len(s) != len(t) {
        return false
    }
    seen := make(map[rune]int)
    for _, r := range s {
        v, _ := seen[r]
        seen[r] = v + 1
    }
    for _, r := range t {
        v, ok := seen[r]
        if !ok {
            return false
        }
        v -= 1
        seen[r] = v
        if v == 0 {
            delete(seen, r)
        }
    }
    return len(seen) == 0
}
