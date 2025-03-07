package events

import "github.com/jitsucom/jitsu/server/logging"

//SegmentProcessor preprocess client integration events
type SegmentProcessor struct {
	usersRecognition Recognition
}

//NewSegmentProcessor returns configured SegmentProcessor
func NewSegmentProcessor(usersRecognition Recognition) *SegmentProcessor {
	return &SegmentProcessor{usersRecognition: usersRecognition}
}

//Preprocess adds src value
//sets user anonymous ID if GDPR
func (sp *SegmentProcessor) Preprocess(event Event, reqContext *RequestContext) {
	event[SrcKey] = "segment_api"

	if !reqContext.CookiesLawCompliant {
		if err := UserAnonymIDPath.Set(event, reqContext.JitsuAnonymousID); err != nil {
			logging.SystemErrorf("Error setting generated Jitsu anonymous ID: %v", err)
		}
	}
	if err := HashedAnonymIDPath.Set(event, reqContext.HashedAnonymousID); err != nil {
		logging.SystemErrorf("Error setting hashed anonymous ID: %v", err)
	}
}

//Postprocess puts event into recognition Service
func (sp *SegmentProcessor) Postprocess(event Event, eventID string, destinationIDs []string) {
	sp.usersRecognition.Event(event, eventID, destinationIDs)
}

//Type returns preprocessor type
func (sp *SegmentProcessor) Type() string {
	return segmentPreprocessorType
}
