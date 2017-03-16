class Raindrops
  SUBSTITUTIONS = {3 => 'Pling',
                   5 => 'Plang',
                   7 => 'Plong'}

  def self.convert(number)
    converted_value = SUBSTITUTIONS.map do |key, value|
      value if number % key == 0
    end.join
    converted_value.empty? ? number.to_s : converted_value
  end
end

module BookKeeping
  VERSION = 3
end
