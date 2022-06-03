
all: gen-OTP gen-secret check-OTP gen-URL gen-QR-code parse-QR-code

gen-OTP: gen-OTP.go
	go build gen-OTP.go
	./gen-OTP --secret "HLYP32J745ZB7LNW" >/tmp/t1.out
	grep "Time based OTP is: ......" /tmp/t1.out 
	@rm /tmp/t1.out 

gen-secret: gen-secret.go
	go build gen-secret.go
	./gen-secret

check-OTP: check-OTP.go gen-OTP
	go build check-OTP.go
	./gen-OTP --raw --secret "HLYP32J745ZB7LNW" | tee ,otp
	./check-OTP --secret "HLYP32J745ZB7LNW" --otp `cat ,otp`
	@rm ,otp

gen-URL: gen-URL.go
	go build gen-URL.go
	./gen-URL

gen-QR-code:
	go build gen-QR-code.go
	./gen-QR-code --URI "otpauth://totp/www.2c-why.com:example%402c-why.com?issuer=www.2c-why.com&secret=HLYP32J745ZB7LNW" --out-img "./a-qr.png"

parse-QR-code:
	go build parse-QR-code.go
	./parse-QR-code  --fn "./a-qr.png" --outURI ,otp
	diff ,otp ../ref/parse-QR-code.txt
	cat ,otp

clean:
	-rm ,otp a-qr.png
	-rm gen-OTP gen-secret check-OTP gen-URL gen-QR-code parse-QR-code

