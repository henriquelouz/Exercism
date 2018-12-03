class Acronym
    def self.abbreviate(name)
        name.split(/\W/).map{ |word| 
            if word[0]
                word[0].upcase
            end
        }.join
    end
end