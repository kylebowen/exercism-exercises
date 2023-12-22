class Pangram

  def self.pangram?(phrase)
    phrase.downcase.gsub(/[^a-z]/, '').chars.uniq.sort.join == ("a".."z").to_a.join
  end
end

module BookKeeping
  VERSION = 4
end
