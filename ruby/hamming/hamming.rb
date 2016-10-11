class Hamming

  def self.compute (a, b)
    return 0 if a == b
    raise ArgumentError if a.length != b.length
    mutations = a.chars.zip(b.chars)
    mutations.delete_if { |a, b| a == b }
    mutations.count
  end
end

module BookKeeping
  VERSION = 3
end
