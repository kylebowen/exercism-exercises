class Hamming

  def self.compute (a, b)
    return 0 if a == b
    raise ArgumentError if a.length != b.length
    find_mutations(a, b).count
  end
end

def find_mutations (a, b)
  pairs = a.chars.zip(b.chars)
  pairs.delete_if { |a, b| a == b }
end

module BookKeeping
  VERSION = 3
end
