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
end

module BookKeeping
  VERSION = 3
end
