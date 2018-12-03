class Isogram
    def self.isogram?(word)
        word.gsub(/\W/,'').each_char {|x|
            if (word.upcase.count x.upcase) > 1
                return false
            end
        }
        return true
    end
end