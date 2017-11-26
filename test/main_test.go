package test

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestKolmogorovSmirnovGaussian(t *testing.T) {
	data1 := []float64{
		+1.8368909586194213e-01, -4.7615301661907383e-01,
		+8.6202161155692203e-01, -1.3616944708707543e+00,
		+4.5502955644433435e-01, -8.4870937993365902e-01,
		-3.3488693896404770e-01, +5.5278334594455014e-01,
		+1.0390906535049560e+00, -1.1176386832652081e+00,
		+1.2606587091208963e+00, +6.6014314104697769e-01,
		-6.7865553542687335e-02, -1.9522119789875436e-01,
		-2.1760635014319193e-01, -3.0310762135174091e-01,
		+2.3045624425105282e-02, +5.1290355848774699e-02,
		+8.2606279021159545e-01, +1.5269766867333727e+00,
	}

	data2 := []float64{
		+4.6691443568470004e-01, -2.0971333838873671e-01,
		+6.2519035708762571e-01, +1.8322726300143696e-01,
		-1.0297675435666211e+00, +9.4922183113102254e-01,
		+3.0706191914670344e-01, +1.3517494209945566e-01,
		+5.1524633552484855e-01, +2.6140632405538267e-01,
		-9.4148577095543373e-01, -1.6233767280382774e-01,
		-1.4605463433152618e-01, -5.3201137680882071e-01,
		+1.6821035946631788e+00, -8.7572934616001730e-01,
		-4.8381505011012110e-01, -7.1200454902742250e-01,
		-1.1742123314568162e+00, -1.9223951753927476e-01,
	}

	rejected, pvalue, statistic := KolmogorovSmirnov(data1, data2, 0.05)

	assert.Equal(rejected, false, t)
	assert.Close(pvalue, 7.7095294467658659e-01, 1e-15, t)
	assert.Close(statistic, 2.0000000000000007e-01, 1e-15, t)
}

func TestKolmogorovSmirnovWeibull(t *testing.T) {
	data1 := []float64{
		8.7461628950465320e-01, 3.2805348316383787e-01,
		9.0760296313544622e+00, 1.1962276336402389e+00,
		1.9189846792253313e+00, 2.3822930800308089e+00,
		1.6806105970365317e+00, 1.0625868855376310e+00,
		9.2440487711596209e-01, 6.1837977701552593e-01,
		8.6942023193092821e-01, 3.7801605351395895e-01,
		1.5874208284094915e+00, 1.2997493985974612e-01,
		3.5976651711297412e+00, 3.9978003354420455e-01,
		8.7393838326219087e-01, 5.8216082843584527e-01,
		1.9633528216664577e+00, 1.6189758084715631e+00,
		2.2221327331258003e-01, 3.2253005358561328e-02,
		1.1601978041124075e+00, 3.6770322439085690e-01,
		1.3194504890926964e-01, 1.1137113942645817e-01,
		2.4645840239524057e+00, 3.2427899202391495e+00,
		1.7729548717711372e+00, 1.2994639392767643e-01,
		2.3192549275168197e+00, 8.6486683662826902e-01,
		4.3022820655631802e-02, 6.2892379974946411e-01,
		3.6834692013026532e-01, 1.1535470546272795e+00,
		3.7614770245443635e-01, 1.8077195177537575e-01,
		4.0014950066798560e+00, 2.8748967103549722e-01,
		1.1201413335031433e-02, 2.9013086245233494e-01,
		1.2713812457411025e+00, 2.3663499233392407e-01,
		2.2708344559604376e+00, 8.0319973959982871e-01,
		9.5855274939282600e-02, 1.2254887941191932e+00,
		1.2455751776154669e+00, 2.0400010671506421e+00,
	}

	data2 := []float64{
		2.3831972409571311e+00, 7.4687483316870418e-01,
		1.4953967997990596e+00, 1.3818064935457834e+00,
		1.0112408625687035e+00, 2.0542956216762875e+00,
		8.9391625059996016e-01, 1.6624096945581250e+00,
		8.7263228664338210e-01, 7.1701401175000756e-01,
		1.8117651960300805e+00, 1.1268211359176015e+00,
		7.2469160298066559e-01, 1.1266309014978020e+00,
		2.0773048597848836e+00, 9.4778307072697676e-01,
		7.6817102463234821e-01, 9.7768980447137177e-01,
		2.8649424647995436e-01, 8.7648378662405579e-01,
		3.8247397215059686e-01, 1.6903887428782887e+00,
		1.6848338500054207e+00, 5.5505268065669600e-01,
		1.1523221867552904e+00, 1.6098193245119850e+00,
		3.2918787364845276e-01, 1.2332741840978334e+00,
		6.4242025954586868e-01, 6.7904313293871110e-01,
		4.2270583693329028e-01, 8.2454091183641387e-01,
		6.4222569780367744e-01, 1.2313746559990257e+00,
		1.3732532939676521e+00, 3.9788978659375451e-01,
		1.1053159845638993e+00, 2.2702828394384506e-01,
		7.6866964666838733e-01, 8.2730797319419624e-01,
		1.7656867546746680e+00, 2.7319723346979158e-01,
		1.0724422049099485e+00, 8.8792504877824741e-01,
		1.1359842068648212e+00, 1.4397906365994388e+00,
		3.8252063532077912e-01, 8.9453092224749142e-01,
		2.9032385714734148e+00, 8.3367741780056803e-01,
	}

	rejected, pvalue, statistic := KolmogorovSmirnov(data1, data2, 0.01)

	assert.Equal(rejected, false, t)
	assert.Close(pvalue, 3.1660846051790786e-02, 1e-15, t)
	assert.Close(statistic, 2.7999999999999997e-01, 1e-15, t)
}
