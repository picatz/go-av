package av

// https://developer.apple.com/documentation/avfoundation
// https://developer.apple.com/documentation/avfoundation/avspeechsynthesisvoice
// https://developer.apple.com/documentation/avfaudio/avspeechsynthesizer

/*
 #cgo CFLAGS: -x objective-c -fmodules -fblocks
 #cgo LDFLAGS: -framework Foundation
 #import <Foundation/Foundation.h>
 #import <AVFoundation/AVFoundation.h>

void textToSpeech(const char *text) {
    NSString *textToSay = [NSString stringWithUTF8String:text];
    AVSpeechSynthesisVoice *voice = [AVSpeechSynthesisVoice voiceWithIdentifier:@"com.apple.ttsbundle.siri_Arthur_en-GB_compact"];
    AVSpeechUtterance *utterance = [AVSpeechUtterance speechUtteranceWithString:textToSay];
    utterance.voice = voice;
    utterance.rate = 0.57;
    utterance.volume = 1.0;
    utterance.postUtteranceDelay = 0.5;
    utterance.pitchMultiplier = 1.0;

    AVSpeechSynthesizer *synthesizer = [[AVSpeechSynthesizer alloc] init];
    [synthesizer speakUtterance:utterance];

    while(synthesizer.speaking);
       sleep(1);
}

void textToSpeechWithVoice(const char *text, const char *voiceIdentifier) {
    NSString *textToSay = [NSString stringWithUTF8String:text];
    NSString *voiceIdentifierString = [NSString stringWithUTF8String:voiceIdentifier];
    AVSpeechSynthesisVoice *voice = [AVSpeechSynthesisVoice voiceWithIdentifier:voiceIdentifierString];
    AVSpeechUtterance *utterance = [AVSpeechUtterance speechUtteranceWithString:textToSay];
    utterance.voice = voice;
    utterance.rate = 0.57;
    utterance.volume = 1.0;
    utterance.postUtteranceDelay = 0.2;
    utterance.pitchMultiplier = 1.0;

    AVSpeechSynthesizer *synthesizer = [[AVSpeechSynthesizer alloc] init];
    [synthesizer speakUtterance:utterance];

    while(synthesizer.speaking);
       sleep(1);
}

void listVoices() {
    NSArray *allVoices = [AVSpeechSynthesisVoice speechVoices];

    for (AVSpeechSynthesisVoice *voice in allVoices) {
        NSLog(@"Voice Name: %@, Identifier: %@, Quality: %ld", voice.name, voice.identifier, (long)voice.quality);
    }
}

AVAudioRecorder *startRecording(const char *filePath) {
    NSError *error;

	// AVAudioSession *session = [AVAudioSession sharedInstance];
    // [session setCategory:AVAudioSessionCategoryRecord error:&error];
    // [session setActive:YES error:&error];

    NSMutableDictionary *settings = [NSMutableDictionary dictionary];
    [settings setValue:@(kAudioFormatMPEG4AAC) forKey:AVFormatIDKey];
    [settings setValue:@(16000.0) forKey:AVSampleRateKey];
    [settings setValue:@(1) forKey:AVNumberOfChannelsKey];

	NSString *filePathString = [NSString stringWithUTF8String:filePath];
	NSURL *url = [NSURL fileURLWithPath:filePathString];
    AVAudioRecorder *recorder = [[AVAudioRecorder alloc] initWithURL:url settings:settings error:&error];

    if (recorder) {
        [recorder prepareToRecord];
        [recorder record];
    } else {
        NSLog(@"%@", [error localizedDescription]);
    }

    return recorder;
}

void stopRecording(AVAudioRecorder *recorder) {
    [recorder stop];
    // AVAudioSession *session = [AVAudioSession sharedInstance];
    // [session setActive:NO error:nil];
}
*/
import "C"
import (
	"context"
	"unsafe"
)

const (
	VoiceSiriFemale = "com.apple.ttsbundle.siri_female_en-GB_premium"
	VoiceSiriMale   = "com.apple.ttsbundle.siri_male_en-GB_premium"

	/*

		TODO: Add more voices

		en-US
		com.apple.voice.compact.ar-001.Maged
		com.apple.voice.compact.bg-BG.Daria
		com.apple.voice.compact.ca-ES.Montserrat
		com.apple.voice.compact.cs-CZ.Zuzana
		com.apple.voice.compact.da-DK.Sara
		com.apple.eloquence.de-DE.Sandy
		com.apple.eloquence.de-DE.Shelley
		com.apple.ttsbundle.siri_Helena_de-DE_compact
		com.apple.eloquence.de-DE.Grandma
		com.apple.eloquence.de-DE.Grandpa
		com.apple.eloquence.de-DE.Eddy
		com.apple.eloquence.de-DE.Reed
		com.apple.voice.compact.de-DE.Anna
		com.apple.ttsbundle.siri_Martin_de-DE_compact
		com.apple.eloquence.de-DE.Rocko
		com.apple.eloquence.de-DE.Flo
		com.apple.voice.compact.el-GR.Melina
		com.apple.ttsbundle.siri_Gordon_en-AU_compact
		com.apple.voice.compact.en-AU.Karen
		com.apple.ttsbundle.siri_Catherine_en-AU_compact
		com.apple.eloquence.en-GB.Rocko
		com.apple.eloquence.en-GB.Shelley
		com.apple.voice.compact.en-GB.Daniel
		com.apple.ttsbundle.siri_Martha_en-GB_compact
		com.apple.eloquence.en-GB.Grandma
		com.apple.eloquence.en-GB.Grandpa
		com.apple.eloquence.en-GB.Flo
		com.apple.eloquence.en-GB.Eddy
		com.apple.eloquence.en-GB.Reed
		com.apple.eloquence.en-GB.Sandy
		com.apple.ttsbundle.siri_Arthur_en-GB_compact
		com.apple.voice.compact.en-IE.Moira
		com.apple.voice.compact.en-IN.Rishi
		com.apple.eloquence.en-US.Flo
		com.apple.speech.synthesis.voice.Bahh
		com.apple.speech.synthesis.voice.Albert
		com.apple.speech.synthesis.voice.Fred
		com.apple.speech.synthesis.voice.Hysterical
		com.apple.speech.synthesis.voice.Organ
		com.apple.speech.synthesis.voice.Cellos
		com.apple.speech.synthesis.voice.Zarvox
		com.apple.eloquence.en-US.Rocko
		com.apple.eloquence.en-US.Shelley
		com.apple.speech.synthesis.voice.Princess
		com.apple.eloquence.en-US.Grandma
		com.apple.eloquence.en-US.Eddy
		com.apple.speech.synthesis.voice.Bells
		com.apple.eloquence.en-US.Grandpa
		com.apple.speech.synthesis.voice.Trinoids
		com.apple.speech.synthesis.voice.Kathy
		com.apple.eloquence.en-US.Reed
		com.apple.speech.synthesis.voice.Boing
		com.apple.speech.synthesis.voice.Whisper
		com.apple.speech.synthesis.voice.GoodNews
		com.apple.speech.synthesis.voice.Deranged
		com.apple.ttsbundle.siri_Nicky_en-US_compact
		com.apple.speech.synthesis.voice.BadNews
		com.apple.ttsbundle.siri_Aaron_en-US_compact
		com.apple.speech.synthesis.voice.Bubbles
		com.apple.voice.compact.en-US.Samantha
		com.apple.eloquence.en-US.Sandy
		com.apple.speech.synthesis.voice.Junior
		com.apple.speech.synthesis.voice.Ralph
		com.apple.voice.compact.en-ZA.Tessa
		com.apple.eloquence.es-ES.Shelley
		com.apple.eloquence.es-ES.Grandma
		com.apple.eloquence.es-ES.Rocko
		com.apple.eloquence.es-ES.Grandpa
		com.apple.eloquence.es-ES.Sandy
		com.apple.voice.compact.es-ES.Monica
		com.apple.eloquence.es-ES.Flo
		com.apple.eloquence.es-ES.Eddy
		com.apple.eloquence.es-ES.Reed
		com.apple.eloquence.es-MX.Rocko
		com.apple.voice.compact.es-MX.Paulina
		com.apple.eloquence.es-MX.Flo
		com.apple.eloquence.es-MX.Sandy
		com.apple.eloquence.es-MX.Eddy
		com.apple.eloquence.es-MX.Shelley
		com.apple.eloquence.es-MX.Grandma
		com.apple.eloquence.es-MX.Reed
		com.apple.eloquence.es-MX.Grandpa
		com.apple.eloquence.fi-FI.Shelley
		com.apple.eloquence.fi-FI.Grandma
		com.apple.eloquence.fi-FI.Grandpa
		com.apple.eloquence.fi-FI.Sandy
		com.apple.voice.compact.fi-FI.Satu
		com.apple.eloquence.fi-FI.Eddy
		com.apple.eloquence.fi-FI.Rocko
		com.apple.eloquence.fi-FI.Reed
		com.apple.eloquence.fi-FI.Flo
		com.apple.eloquence.fr-CA.Shelley
		com.apple.eloquence.fr-CA.Grandma
		com.apple.eloquence.fr-CA.Grandpa
		com.apple.eloquence.fr-CA.Rocko
		com.apple.eloquence.fr-CA.Eddy
		com.apple.eloquence.fr-CA.Reed
		com.apple.voice.compact.fr-CA.Amelie
		com.apple.eloquence.fr-CA.Flo
		com.apple.eloquence.fr-CA.Sandy
		com.apple.eloquence.fr-FR.Grandma
		com.apple.eloquence.fr-FR.Flo
		com.apple.eloquence.fr-FR.Rocko
		com.apple.eloquence.fr-FR.Grandpa
		com.apple.eloquence.fr-FR.Sandy
		com.apple.eloquence.fr-FR.Eddy
		com.apple.voice.compact.fr-FR.Thomas
		com.apple.ttsbundle.siri_Dan_fr-FR_compact
		com.apple.eloquence.fr-FR.Jacques
		com.apple.ttsbundle.siri_Marie_fr-FR_compact
		com.apple.eloquence.fr-FR.Shelley
		com.apple.voice.compact.he-IL.Carmit
		com.apple.voice.compact.hi-IN.Lekha
		com.apple.voice.compact.hr-HR.Lana
		com.apple.voice.compact.hu-HU.Mariska
		com.apple.voice.compact.id-ID.Damayanti
		com.apple.eloquence.it-IT.Eddy
		com.apple.eloquence.it-IT.Sandy
		com.apple.eloquence.it-IT.Reed
		com.apple.eloquence.it-IT.Shelley
		com.apple.eloquence.it-IT.Grandma
		com.apple.eloquence.it-IT.Grandpa
		com.apple.eloquence.it-IT.Flo
		com.apple.eloquence.it-IT.Rocko
		com.apple.voice.compact.it-IT.Alice
		com.apple.voice.compact.ja-JP.Kyoko
		com.apple.ttsbundle.siri_Hattori_ja-JP_compact
		com.apple.ttsbundle.siri_O-Ren_ja-JP_compact
		com.apple.voice.compact.ko-KR.Yuna
		com.apple.voice.compact.ms-MY.Amira
		com.apple.voice.compact.nb-NO.Nora
		com.apple.voice.compact.nl-BE.Ellen
		com.apple.voice.compact.nl-NL.Xander
		com.apple.voice.compact.pl-PL.Zosia
		com.apple.eloquence.pt-BR.Reed
		com.apple.voice.compact.pt-BR.Luciana
		com.apple.eloquence.pt-BR.Shelley
		com.apple.eloquence.pt-BR.Grandma
		com.apple.eloquence.pt-BR.Grandpa
		com.apple.eloquence.pt-BR.Rocko
		com.apple.eloquence.pt-BR.Flo
		com.apple.eloquence.pt-BR.Sandy
		com.apple.eloquence.pt-BR.Eddy
		com.apple.voice.compact.pt-PT.Joana
		com.apple.voice.compact.ro-RO.Ioana
		com.apple.voice.compact.ru-RU.Milena
		com.apple.voice.compact.sk-SK.Laura
		com.apple.voice.compact.sv-SE.Alva
		com.apple.voice.compact.th-TH.Kanya
		com.apple.voice.compact.tr-TR.Yelda
		com.apple.voice.compact.uk-UA.Lesya
		com.apple.voice.compact.vi-VN.Linh
		com.apple.ttsbundle.siri_Yu-Shu_zh-CN_compact
		com.apple.ttsbundle.siri_Li-Mu_zh-CN_compact
		com.apple.voice.compact.zh-CN.Tingting
		com.apple.voice.compact.zh-HK.Sinji
		com.apple.voice.compact.zh-TW.Meijia





		com.apple.ttsbundle.Tessa-compact
		com.apple.ttsbundle.Karen-compact
		com.apple.speech.voice.Alex
		com.apple.ttsbundle.siri_female_en-GB_premium
		com.apple.ttsbundle.Samantha-compact
		com.apple.ttsbundle.siri_female_en-US_premium
		com.apple.eloquence.en-US.Rocko
		com.apple.eloquence.en-US.Shelley
		com.apple.speech.synthesis.voice.GoodNews <- it sings?
		com.apple.speech.synthesis.voice.Bubbles
		com.apple.ttsbundle.siri_Gordon_en-AU_compact

		com.apple.ttsbundle.siri_male_en-GB_compact
		com.apple.ttsbundle.siri_female_en-GB_compact
		com.apple.voice.premium.en-GB.Serena
		com.apple.voice.enhanced.en-GB.Serena
		com.apple.voice.enhanced.en-GB.Daniel

		com.apple.ttsbundle.siri_Martha_en-GB_compact
		com.apple.ttsbundle.siri_female_en-GB_premium
		com.apple.ttsbundle.siri_Arthur_en-GB_premium
	*/
)

// TextToSpeech speaks the text with the default voice (Siri female).
func TextToSpeech(text string) {
	C.textToSpeech(C.CString(text))
}

// TextToSpeechWithVoice speaks the text with the given voice.
func TextToSpeechWithVoice(text string, voice string) {
	C.textToSpeechWithVoice(C.CString(text), C.CString(voice))
}

// PrintVoices prints the available voices to stdout.
func PrintVoices() {
	C.listVoices()
}

// RecordAudioToFile records audio to a file until the context is done.
//
// File format is determined by the file extension, which must be ".m4a" to work.
func RecordAudioToFile(ctx context.Context, path string) {
	cfilePath := C.CString(path)
	defer C.free(unsafe.Pointer(cfilePath))

	// Start recording
	recorder := C.startRecording(cfilePath)
	defer C.free(unsafe.Pointer(recorder))

	// Record for duration until context is done.
	<-ctx.Done()

	// Stop recording
	C.stopRecording(recorder)
}
