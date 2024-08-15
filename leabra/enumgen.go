// Code generated by "core generate"; DO NOT EDIT.

package leabra

import (
	"cogentcore.org/core/enums"
)

var _ActNoiseTypeValues = []ActNoiseType{0, 1, 2, 3, 4}

// ActNoiseTypeN is the highest valid value for type ActNoiseType, plus one.
const ActNoiseTypeN ActNoiseType = 5

var _ActNoiseTypeValueMap = map[string]ActNoiseType{`NoNoise`: 0, `VmNoise`: 1, `GeNoise`: 2, `ActNoise`: 3, `GeMultNoise`: 4}

var _ActNoiseTypeDescMap = map[ActNoiseType]string{0: `NoNoise means no noise added`, 1: `VmNoise means noise is added to the membrane potential. IMPORTANT: this should NOT be used for rate-code (NXX1) activations, because they do not depend directly on the vm -- this then has no effect`, 2: `GeNoise means noise is added to the excitatory conductance (Ge). This should be used for rate coded activations (NXX1)`, 3: `ActNoise means noise is added to the final rate code activation`, 4: `GeMultNoise means that noise is multiplicative on the Ge excitatory conductance values`}

var _ActNoiseTypeMap = map[ActNoiseType]string{0: `NoNoise`, 1: `VmNoise`, 2: `GeNoise`, 3: `ActNoise`, 4: `GeMultNoise`}

// String returns the string representation of this ActNoiseType value.
func (i ActNoiseType) String() string { return enums.String(i, _ActNoiseTypeMap) }

// SetString sets the ActNoiseType value from its string representation,
// and returns an error if the string is invalid.
func (i *ActNoiseType) SetString(s string) error {
	return enums.SetString(i, s, _ActNoiseTypeValueMap, "ActNoiseType")
}

// Int64 returns the ActNoiseType value as an int64.
func (i ActNoiseType) Int64() int64 { return int64(i) }

// SetInt64 sets the ActNoiseType value from an int64.
func (i *ActNoiseType) SetInt64(in int64) { *i = ActNoiseType(in) }

// Desc returns the description of the ActNoiseType value.
func (i ActNoiseType) Desc() string { return enums.Desc(i, _ActNoiseTypeDescMap) }

// ActNoiseTypeValues returns all possible values for the type ActNoiseType.
func ActNoiseTypeValues() []ActNoiseType { return _ActNoiseTypeValues }

// Values returns all possible values for the type ActNoiseType.
func (i ActNoiseType) Values() []enums.Enum { return enums.Values(_ActNoiseTypeValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i ActNoiseType) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *ActNoiseType) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "ActNoiseType")
}

var _LayerTypesValues = []LayerTypes{0, 1, 2, 3, 4, 5, 6, 7, 8}

// LayerTypesN is the highest valid value for type LayerTypes, plus one.
const LayerTypesN LayerTypes = 9

var _LayerTypesValueMap = map[string]LayerTypes{`SuperLayer`: 0, `InputLayer`: 1, `TargetLayer`: 2, `CompareLayer`: 3, `CTLayer`: 4, `PulvinarLayer`: 5, `TRNLayer`: 6, `PTMaintLayer`: 7, `PTPredLayer`: 8}

var _LayerTypesDescMap = map[LayerTypes]string{0: `Super is a superficial cortical layer (lamina 2-3-4) which does not receive direct input or targets. In more generic models, it should be used as a Hidden layer, and maps onto the Hidden type in LayerTypes.`, 1: `Input is a layer that receives direct external input in its Ext inputs. Biologically, it can be a primary sensory layer, or a thalamic layer.`, 2: `Target is a layer that receives direct external target inputs used for driving plus-phase learning. Simple target layers are generally not used in more biological models, which instead use predictive learning via Pulvinar or related mechanisms.`, 3: `Compare is a layer that receives external comparison inputs, which drive statistics but do NOT drive activation or learning directly. It is rarely used in axon.`, 4: `CT are layer 6 corticothalamic projecting neurons, which drive &#34;top down&#34; predictions in Pulvinar layers. They maintain information over time via stronger NMDA channels and use maintained prior state information to generate predictions about current states forming on Super layers that then drive PT (5IB) bursting activity, which are the plus-phase drivers of Pulvinar activity.`, 5: `Pulvinar are thalamic relay cell neurons in the higher-order Pulvinar nucleus of the thalamus, and functionally isomorphic neurons in the MD thalamus, and potentially other areas. These cells alternately reflect predictions driven by CT pathways, and actual outcomes driven by 5IB Burst activity from corresponding PT or Super layer neurons that provide strong driving inputs.`, 6: `TRNLayer is thalamic reticular nucleus layer for inhibitory competition within the thalamus.`, 7: `PTMaintLayer implements the subset of pyramidal tract (PT) layer 5 intrinsic bursting (5IB) deep neurons that exhibit robust, stable maintenance of activity over the duration of a goal engaged window, modulated by basal ganglia (BG) disinhibitory gating, supported by strong MaintNMDA channels and recurrent excitation. The lateral PTSelfMaint pathway uses MaintG to drive GMaintRaw input that feeds into the stronger, longer MaintNMDA channels, and the ThalToPT ModulatoryG pathway from BGThalamus multiplicatively modulates the strength of other inputs, such that only at the time of BG gating are these strong enough to drive sustained active maintenance. Use Act.Dend.ModGain to parameterize.`, 8: `PTPredLayer implements the subset of pyramidal tract (PT) layer 5 intrinsic bursting (5IB) deep neurons that combine modulatory input from PTMaintLayer sustained maintenance and CTLayer dynamic predictive learning that helps to predict state changes during the period of active goal maintenance. This layer provides the primary input to VSPatch US-timing prediction layers, and other layers that require predictive dynamic`}

var _LayerTypesMap = map[LayerTypes]string{0: `SuperLayer`, 1: `InputLayer`, 2: `TargetLayer`, 3: `CompareLayer`, 4: `CTLayer`, 5: `PulvinarLayer`, 6: `TRNLayer`, 7: `PTMaintLayer`, 8: `PTPredLayer`}

// String returns the string representation of this LayerTypes value.
func (i LayerTypes) String() string { return enums.String(i, _LayerTypesMap) }

// SetString sets the LayerTypes value from its string representation,
// and returns an error if the string is invalid.
func (i *LayerTypes) SetString(s string) error {
	return enums.SetString(i, s, _LayerTypesValueMap, "LayerTypes")
}

// Int64 returns the LayerTypes value as an int64.
func (i LayerTypes) Int64() int64 { return int64(i) }

// SetInt64 sets the LayerTypes value from an int64.
func (i *LayerTypes) SetInt64(in int64) { *i = LayerTypes(in) }

// Desc returns the description of the LayerTypes value.
func (i LayerTypes) Desc() string { return enums.Desc(i, _LayerTypesDescMap) }

// LayerTypesValues returns all possible values for the type LayerTypes.
func LayerTypesValues() []LayerTypes { return _LayerTypesValues }

// Values returns all possible values for the type LayerTypes.
func (i LayerTypes) Values() []enums.Enum { return enums.Values(_LayerTypesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i LayerTypes) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *LayerTypes) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "LayerTypes")
}

var _NeurFlagsValues = []NeurFlags{0, 1, 2, 3}

// NeurFlagsN is the highest valid value for type NeurFlags, plus one.
const NeurFlagsN NeurFlags = 4

var _NeurFlagsValueMap = map[string]NeurFlags{`NeurOff`: 0, `NeurHasExt`: 1, `NeurHasTarg`: 2, `NeurHasCmpr`: 3}

var _NeurFlagsDescMap = map[NeurFlags]string{0: `NeurOff flag indicates that this neuron has been turned off (i.e., lesioned)`, 1: `NeurHasExt means the neuron has external input in its Ext field`, 2: `NeurHasTarg means the neuron has external target input in its Targ field`, 3: `NeurHasCmpr means the neuron has external comparison input in its Targ field -- used for computing comparison statistics but does not drive neural activity ever`}

var _NeurFlagsMap = map[NeurFlags]string{0: `NeurOff`, 1: `NeurHasExt`, 2: `NeurHasTarg`, 3: `NeurHasCmpr`}

// String returns the string representation of this NeurFlags value.
func (i NeurFlags) String() string { return enums.BitFlagString(i, _NeurFlagsValues) }

// BitIndexString returns the string representation of this NeurFlags value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i NeurFlags) BitIndexString() string { return enums.String(i, _NeurFlagsMap) }

// SetString sets the NeurFlags value from its string representation,
// and returns an error if the string is invalid.
func (i *NeurFlags) SetString(s string) error { *i = 0; return i.SetStringOr(s) }

// SetStringOr sets the NeurFlags value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *NeurFlags) SetStringOr(s string) error {
	return enums.SetStringOr(i, s, _NeurFlagsValueMap, "NeurFlags")
}

// Int64 returns the NeurFlags value as an int64.
func (i NeurFlags) Int64() int64 { return int64(i) }

// SetInt64 sets the NeurFlags value from an int64.
func (i *NeurFlags) SetInt64(in int64) { *i = NeurFlags(in) }

// Desc returns the description of the NeurFlags value.
func (i NeurFlags) Desc() string { return enums.Desc(i, _NeurFlagsDescMap) }

// NeurFlagsValues returns all possible values for the type NeurFlags.
func NeurFlagsValues() []NeurFlags { return _NeurFlagsValues }

// Values returns all possible values for the type NeurFlags.
func (i NeurFlags) Values() []enums.Enum { return enums.Values(_NeurFlagsValues) }

// HasFlag returns whether these bit flags have the given bit flag set.
func (i NeurFlags) HasFlag(f enums.BitFlag) bool { return enums.HasFlag((*int64)(&i), f) }

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *NeurFlags) SetFlag(on bool, f ...enums.BitFlag) { enums.SetFlag((*int64)(i), on, f...) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i NeurFlags) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *NeurFlags) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "NeurFlags")
}

var _PathTypesValues = []PathTypes{0, 1, 2, 3, 4}

// PathTypesN is the highest valid value for type PathTypes, plus one.
const PathTypesN PathTypes = 5

var _PathTypesValueMap = map[string]PathTypes{`ForwardPath`: 0, `BackPath`: 1, `LateralPath`: 2, `InhibPath`: 3, `CTCtxtPath`: 4}

var _PathTypesDescMap = map[PathTypes]string{0: `Forward is a feedforward, bottom-up pathway from sensory inputs to higher layers`, 1: `Back is a feedback, top-down pathway from higher layers back to lower layers`, 2: `Lateral is a lateral pathway within the same layer / area`, 3: `Inhib is an inhibitory pathway that drives inhibitory synaptic conductances instead of the default excitatory ones.`, 4: `CTCtxt are pathways from Superficial layers to CT layers that send Burst activations drive updating of CtxtGe excitatory conductance, at end of plus (51B Bursting) phase. Biologically, this pathway comes from the PT layer 5IB neurons, but it is simpler to use the Super neurons directly, and PT are optional for most network types. These pathways also use a special learning rule that takes into account the temporal delays in the activation states. Can also add self context from CT for deeper temporal context.`}

var _PathTypesMap = map[PathTypes]string{0: `ForwardPath`, 1: `BackPath`, 2: `LateralPath`, 3: `InhibPath`, 4: `CTCtxtPath`}

// String returns the string representation of this PathTypes value.
func (i PathTypes) String() string { return enums.String(i, _PathTypesMap) }

// SetString sets the PathTypes value from its string representation,
// and returns an error if the string is invalid.
func (i *PathTypes) SetString(s string) error {
	return enums.SetString(i, s, _PathTypesValueMap, "PathTypes")
}

// Int64 returns the PathTypes value as an int64.
func (i PathTypes) Int64() int64 { return int64(i) }

// SetInt64 sets the PathTypes value from an int64.
func (i *PathTypes) SetInt64(in int64) { *i = PathTypes(in) }

// Desc returns the description of the PathTypes value.
func (i PathTypes) Desc() string { return enums.Desc(i, _PathTypesDescMap) }

// PathTypesValues returns all possible values for the type PathTypes.
func PathTypesValues() []PathTypes { return _PathTypesValues }

// Values returns all possible values for the type PathTypes.
func (i PathTypes) Values() []enums.Enum { return enums.Values(_PathTypesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i PathTypes) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *PathTypes) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "PathTypes")
}

var _QuartersValues = []Quarters{0, 1, 2, 3}

// QuartersN is the highest valid value for type Quarters, plus one.
const QuartersN Quarters = 4

var _QuartersValueMap = map[string]Quarters{`Q1`: 0, `Q2`: 1, `Q3`: 2, `Q4`: 3}

var _QuartersDescMap = map[Quarters]string{0: `Q1 is the first quarter, which, due to 0-based indexing, shows up as Quarter = 0 in timer`, 1: ``, 2: ``, 3: ``}

var _QuartersMap = map[Quarters]string{0: `Q1`, 1: `Q2`, 2: `Q3`, 3: `Q4`}

// String returns the string representation of this Quarters value.
func (i Quarters) String() string { return enums.BitFlagString(i, _QuartersValues) }

// BitIndexString returns the string representation of this Quarters value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i Quarters) BitIndexString() string { return enums.String(i, _QuartersMap) }

// SetString sets the Quarters value from its string representation,
// and returns an error if the string is invalid.
func (i *Quarters) SetString(s string) error { *i = 0; return i.SetStringOr(s) }

// SetStringOr sets the Quarters value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *Quarters) SetStringOr(s string) error {
	return enums.SetStringOr(i, s, _QuartersValueMap, "Quarters")
}

// Int64 returns the Quarters value as an int64.
func (i Quarters) Int64() int64 { return int64(i) }

// SetInt64 sets the Quarters value from an int64.
func (i *Quarters) SetInt64(in int64) { *i = Quarters(in) }

// Desc returns the description of the Quarters value.
func (i Quarters) Desc() string { return enums.Desc(i, _QuartersDescMap) }

// QuartersValues returns all possible values for the type Quarters.
func QuartersValues() []Quarters { return _QuartersValues }

// Values returns all possible values for the type Quarters.
func (i Quarters) Values() []enums.Enum { return enums.Values(_QuartersValues) }

// HasFlag returns whether these bit flags have the given bit flag set.
func (i Quarters) HasFlag(f enums.BitFlag) bool { return enums.HasFlag((*int64)(&i), f) }

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *Quarters) SetFlag(on bool, f ...enums.BitFlag) { enums.SetFlag((*int64)(i), on, f...) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Quarters) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Quarters) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Quarters") }

var _TimeScalesValues = []TimeScales{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

// TimeScalesN is the highest valid value for type TimeScales, plus one.
const TimeScalesN TimeScales = 18

var _TimeScalesValueMap = map[string]TimeScales{`Cycle`: 0, `FastSpike`: 1, `Quarter`: 2, `Phase`: 3, `BetaCycle`: 4, `AlphaCycle`: 5, `ThetaCycle`: 6, `Event`: 7, `Trial`: 8, `Tick`: 9, `Sequence`: 10, `Condition`: 11, `Block`: 12, `Epoch`: 13, `Run`: 14, `Expt`: 15, `Scene`: 16, `Episode`: 17}

var _TimeScalesDescMap = map[TimeScales]string{0: `Cycle is the finest time scale -- typically 1 msec -- a single activation update.`, 1: `FastSpike is typically 10 cycles = 10 msec (100hz) = the fastest spiking time generally observed in the brain. This can be useful for visualizing updates at a granularity in between Cycle and Quarter.`, 2: `Quarter is typically 25 cycles = 25 msec (40hz) = 1/4 of the 100 msec alpha trial This is also the GammaCycle (gamma = 40hz), but we use Quarter functionally by virtue of there being 4 per AlphaCycle.`, 3: `Phase is either Minus or Plus phase -- Minus = first 3 quarters, Plus = last quarter`, 4: `BetaCycle is typically 50 cycles = 50 msec (20 hz) = one beta-frequency cycle. Gating in the basal ganglia and associated updating in prefrontal cortex occurs at this frequency.`, 5: `AlphaCycle is typically 100 cycles = 100 msec (10 hz) = one alpha-frequency cycle, which is the fundamental unit of learning in posterior cortex.`, 6: `ThetaCycle is typically 200 cycles = 200 msec (5 hz) = two alpha-frequency cycles. This is the modal duration of a saccade, the update frequency of medial temporal lobe episodic memory, and the minimal predictive learning cycle (perceive an Alpha 1, predict on 2).`, 7: `Event is the smallest unit of naturalistic experience that coheres unto itself (e.g., something that could be described in a sentence). Typically this is on the time scale of a few seconds: e.g., reaching for something, catching a ball.`, 8: `Trial is one unit of behavior in an experiment -- it is typically environmentally defined instead of endogenously defined in terms of basic brain rhythms. In the minimal case it could be one AlphaCycle, but could be multiple, and could encompass multiple Events (e.g., one event is fixation, next is stimulus, last is response)`, 9: `Tick is one step in a sequence -- often it is useful to have Trial count up throughout the entire Epoch but also include a Tick to count trials within a Sequence`, 10: `Sequence is a sequential group of Trials (not always needed).`, 11: `Condition is a collection of Blocks that share the same set of parameters. This is intermediate between Block and Run levels.`, 12: `Block is a collection of Trials, Sequences or Events, often used in experiments when conditions are varied across blocks.`, 13: `Epoch is used in two different contexts. In machine learning, it represents a collection of Trials, Sequences or Events that constitute a &#34;representative sample&#34; of the environment. In the simplest case, it is the entire collection of Trials used for training. In electrophysiology, it is a timing window used for organizing the analysis of electrode data.`, 14: `Run is a complete run of a model / subject, from training to testing, etc. Often multiple runs are done in an Expt to obtain statistics over initial random weights etc.`, 15: `Expt is an entire experiment -- multiple Runs through a given protocol / set of parameters.`, 16: `Scene is a sequence of events that constitutes the next larger-scale coherent unit of naturalistic experience corresponding e.g., to a scene in a movie. Typically consists of events that all take place in one location over e.g., a minute or so. This could be a paragraph or a page or so in a book.`, 17: `Episode is a sequence of scenes that constitutes the next larger-scale unit of naturalistic experience e.g., going to the grocery store or eating at a restaurant, attending a wedding or other &#34;event&#34;. This could be a chapter in a book.`}

var _TimeScalesMap = map[TimeScales]string{0: `Cycle`, 1: `FastSpike`, 2: `Quarter`, 3: `Phase`, 4: `BetaCycle`, 5: `AlphaCycle`, 6: `ThetaCycle`, 7: `Event`, 8: `Trial`, 9: `Tick`, 10: `Sequence`, 11: `Condition`, 12: `Block`, 13: `Epoch`, 14: `Run`, 15: `Expt`, 16: `Scene`, 17: `Episode`}

// String returns the string representation of this TimeScales value.
func (i TimeScales) String() string { return enums.String(i, _TimeScalesMap) }

// SetString sets the TimeScales value from its string representation,
// and returns an error if the string is invalid.
func (i *TimeScales) SetString(s string) error {
	return enums.SetString(i, s, _TimeScalesValueMap, "TimeScales")
}

// Int64 returns the TimeScales value as an int64.
func (i TimeScales) Int64() int64 { return int64(i) }

// SetInt64 sets the TimeScales value from an int64.
func (i *TimeScales) SetInt64(in int64) { *i = TimeScales(in) }

// Desc returns the description of the TimeScales value.
func (i TimeScales) Desc() string { return enums.Desc(i, _TimeScalesDescMap) }

// TimeScalesValues returns all possible values for the type TimeScales.
func TimeScalesValues() []TimeScales { return _TimeScalesValues }

// Values returns all possible values for the type TimeScales.
func (i TimeScales) Values() []enums.Enum { return enums.Values(_TimeScalesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i TimeScales) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *TimeScales) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "TimeScales")
}
