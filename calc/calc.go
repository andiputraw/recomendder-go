package calc

import "math"

func CalcTf(termFreq uint, totalTerm uint) float64 {
	return float64(termFreq) / float64(totalTerm)
}

func CalcIdf(termFreqDoc, total_doc uint) float64 {
	return math.Log10(float64(total_doc) / float64(termFreqDoc))
}

func CalcTfIdf(termFreq, termFreqDoc, total_term, total_doc uint) float64 {
	return CalcTf(termFreq, total_term) * CalcIdf(termFreqDoc, total_doc)
}