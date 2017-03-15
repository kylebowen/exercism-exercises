class Raindrops
  SUBSTITUTIONS = {3 => 'Pling',
                   5 => 'Plang',
                   7 => 'Plong'}

  def self.convert(number)
    converted_value = ''
    SUBSTITUTIONS.each do |key, value|
      converted_value << value if number % key == 0
    end
    converted_value.empty? ? number.to_s : converted_value
  end
  # Shameless green
  # def self.convert(number)
  #   drops = ''
  #   if number % 3 == 0
  #     drops << "Pling"
  #   end
  #   if number % 5 == 0
  #     drops << "Plang"
  #   end
  #   if number % 7 == 0
  #     drops << "Plong"
  #   end
  #   if number % 3 != 0 && number % 5 != 0 && number % 7 != 0
  #     drops << number.to_s
  #   end
  #   drops
  # end
end

module BookKeeping
  VERSION = 3
end
