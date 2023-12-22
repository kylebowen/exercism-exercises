class Sieve
  attr_reader   :limit
  attr_accessor :range

  def initialize limit
    @limit = limit
    @range = (2..limit).to_a
  end

  def primes
  end
end
