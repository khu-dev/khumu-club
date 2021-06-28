package data

func String2StringPtr(s string, nilIfEmpty bool) *string {
    if len(s) == 0 {
        if nilIfEmpty {
            return nil
        } else{
            tmp := ""
            return &tmp
        }

    } else{
        tmp := s
        return &tmp
    }
}