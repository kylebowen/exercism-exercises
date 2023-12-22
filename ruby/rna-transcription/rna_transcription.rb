class Complement

  def self.of_dna dna_strand
    return '' if dna_strand.match(/[^CGTA]/)
    dna_strand.tr('CGTA', 'GCAU')
  end
end

module BookKeeping
  VERSION = 4
end
