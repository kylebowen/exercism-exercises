class Raindrops

  # Shameless green
  def self.convert(number)
    drops = ''
    if number % 3 == 0
      drops << "Pling"
    end
    if number % 5 == 0
      drops << "Plang"
    end
    if number % 7 == 0
      drops << "Plong"
    end
    if number % 3 != 0 && number % 5 != 0 && number % 7 != 0
      drops << number.to_s
    end
    drops
  end
end

module BookKeeping
  VERSION = 3
end
