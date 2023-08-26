package encryption_test

import (
	"fmt"

	"testing"

	"github.com/mandarinkb/go-example-lib/util/encryption"
	"github.com/mandarinkb/go-example-lib/util/logg"
)

func TestRandom32Char(t *testing.T) {

	random32char := encryption.RandomString(32)
	logg.Printlogger(" ***** TEST random char32 *****", random32char)
}

func TestRandom6Number(t *testing.T) {

	random6number := encryption.RandomNumber(6)
	logg.Printlogger(" ***** TEST random number6 *****", random6number)
}

func TestEncryptParamsValue(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TEST Encryption",
			args: args{
				param: "123",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want = encryption.EncryptParamsValue(tt.args.param); tt.want == "" {
				logg.PrintloggerHasHeader(" ***** EncryptParamsValue() ERROR *****", "", tt.want)
			} else {
				logg.PrintloggerHasHeader(" ***** EncryptParamsValue() SUCCESS *****", "", tt.want)
			}
		})
	}
}

func TestDecryptParamsValue(t *testing.T) {
	type args struct {
		DecryptString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TEST configAllChecking",
			args: args{
				DecryptString: "372c6c57c4e144f18f5637356a113cd8b4928a8760bb648319162622601f59a8a45ca9e6a714a3fda37feff78dce130d49b87c4f41139cfcdd63c1d246bbf3c9755329a550fda087a7349ffc9153b6d3ebce4a05cd8ec8d0720da1a5e30d9c2b943a464fe479259e14a758fded97a216b2c8e684138106a57be0bb62047a96c5c5141229502060ef3b2b600ef6dfdb190ba88c131e3c1f886dff55ef7fb0a18ab45df9ab39a55abde43d0a7418b825eaf27d000385d0a93ca221567f42d5b684523d2bd2e3e44d7d83f371922877dc35b8fde601f35af44f86aaef5280cc5dd11d7c0b369d9e9a1bd84d5f340828bfce33dc93b5c98873776395d5f7517cdadb0e1888f2b8ac0a5e8ffd60b0465b863db8bbd6402c3391334ffa5625b214a3f46ed54277bc74ad8c9f2784bd2a55f504caacb3ab3886643f954e4633f209631fbff9f8e286f9cd794f5d2af543c3cb3ac0625c94e99f5b8947a74280f2e414893aa73fd6449559466556b69f758e70e30a52cb0a4cb5a4d95a0059df32f94cc297a93c28766ff3b198803ea0d089c453df4999e7cdee3890286a48ddcf4bf3a651a9e247d44e6debdbcab7194c14a90323b8747319f484c7f907cb45bafb368758fb0964485780500d",
			},
			want: "",
		},
		{
			name: "TEST configurationChecking",
			args: args{
				DecryptString: "2de2e2113e47607c8f50191ec301165eeeb46d3112f33b76ffaf5837f8059f3dff",
			},
			want: "",
		},
		{
			name: "TEST databaseChecking",
			args: args{
				DecryptString: "77c67c0cefb229b4ae340b5ed211bba818131a3423638ac943380099c1e3ec826398dfeca26592ed6f33e0f77282375bb4f30602a609abd2d9e11e1e3fea7a3cf253e88b67724dd11c9aa6ffff2cd4f744dd46cda7331f7704a495fca59ce6b933506283a3afd410f18fd463",
			},
			want: "",
		},
		{
			name: "TEST messageFileChecking",
			args: args{
				DecryptString: "78f27dbfa0d681ea4d9ae0244710f31d3ece16f0d8a67eac3b1af1432e5534c4917c3bb88b3c8cdb2015e9b307b7f3a1d753ccce36de30351a2edc16deb9b6aa266f323b371db74376",
			},
			want: "",
		},
		{
			name: "TEST redisChecking",
			args: args{
				DecryptString: "bc6227e56d640492164588305605469fad8ef413cc44b9cd8471ba098b6db15f1f791151793b6b8cde295056697bbf094b92ea808ec35de1c445623f6f573c255b9dc7cf4b6ca7c705a5acad6311b428",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encryption.DecryptParamsValue(tt.args.DecryptString)
			if got != tt.want {
				logg.PrintloggerHasHeader(" ***** DecryptParamsValue() ERROR *****", "", got)
			} else {
				logg.PrintloggerHasHeader(" ***** DecryptParamsValue() SUCCESS *****", "", got)
			}

		})
	}
}

func TestCryptoEncrypt(t *testing.T) {
	data := "ABC.2"
	passPhase := "We@reTheBe$t^^"
	encrypted, err := encryption.CryptoJsAesEncrypt(passPhase, data)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("encrypted: " + encrypted)

	decrypted, err := encryption.CryptoJsAesDecrypt(passPhase, encrypted)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("decrypted: ")
	fmt.Println(decrypted)
}

func TestEncodeBase64(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				value: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.args.value = encryption.RandomString(60)
			encode := encryption.EncodeBase64(tt.args.value)
			decode, err := encryption.DecodeBase64(encode)
			if err != nil {
				fmt.Println("decode eror: ", err)
			}
			logg.Printlogger("random value for encode char60 *****", tt.args.value)
			fmt.Println("encode: ", encode)
			fmt.Println("decode: ", decode)
		})
	}
}
