class Complement

  def self.of_dna dna_strand
    rna_strand = dna_strand.chars.map do |nucleotide|
      substitute_complement(nucleotide)
    end.join
    rna_strand.length == dna_strand.length ? rna_strand : ''
  end

  def self.substitute_complement(nucleotide)
    case nucleotide
    when 'C'
      'G'
    when 'G'
      'C'
    when 'T'
      'A'
    when 'A'
      'U'
    end
  end
end

module BookKeeping
  VERSION = 4
end
