class Squares

  def initialize num
    @num = num
  end

  def square_of_sum
    ((@num * (@num + 1))/2) ** 2
  end

  def sum_of_squares
    (@num * (@num + 1) * ((2 * @num) + 1)) / 6
  end

  def difference
    square_of_sum - sum_of_squares
  end
end

module BookKeeping
  VERSION = 3
end
