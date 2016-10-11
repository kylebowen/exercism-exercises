class Hamming

  def self.compute (strand1, strand2)
    raise ArgumentError if strand1.length != strand2.length
    mutations = strand1.chars.zip(strand2.chars)
    mutations.delete_if { |nucleotide1, nucleotide2|
      nucleotide1 == nucleotide2 }
    mutations.count
  end
end

module BookKeeping
  VERSION = 3
end
