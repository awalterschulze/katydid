//  Copyright 2016 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package mem_test

import (
	"math/rand"
	"testing"
	"time"

	gogoproto "github.com/gogo/protobuf/proto"
	katyproto "github.com/katydid/katydid/parser/proto"
	"github.com/katydid/katydid/relapse/mem/testproto"
	rparser "github.com/katydid/katydid/relapse/parser"
)

func TestLargeGrammar(t *testing.T) {

	const long = `(.uint(8):== "422"&((.uint(10):== "02"&(((((((((((((((((((((((((((((((((((((((((((((((((((((((((((.uint(14):->contains($uint,[]uint{uint(12871)})&.uint(12):== uint(40111))|(.uint(14):->contains($uint,[]uint{uint(12891)})&.uint(12):== uint(50105)))|(.uint(14):->contains($uint,[]uint{uint(14841)})&.uint(12):== uint(50101)))|(.uint(14):->contains($uint,[]uint{uint(14881)})&.uint(12):== uint(40110)))|(.uint(14):->contains($uint,[]uint{uint(23610)})&.uint(12):== uint(7006)))|(.uint(14):->contains($uint,[]uint{uint(26133)})&.uint(12):== uint(20102)))|(.uint(14):->contains($uint,[]uint{uint(28602)})&.uint(12):== uint(7003)))|(.uint(14):->contains($uint,[]uint{uint(29440)})&.uint(12):== uint(10007)))|(.uint(14):->contains($uint,[]uint{uint(31702)})&.uint(12):== uint(1012)))|(.uint(14):->contains($uint,[]uint{uint(31742)})&.uint(12):== uint(1011)))|(.uint(14):->contains($uint,[]uint{uint(32261)})&.uint(12):== uint(20111)))|(.uint(14):->contains($uint,[]uint{uint(33375)})&.uint(12):== uint(3152)))|(.uint(14):->contains($uint,[]uint{uint(33723)})&.uint(12):== uint(3010)))|(.uint(14):->contains($uint,[]uint{uint(37192)})&.uint(12):== uint(2107)))|(.uint(14):->contains($uint,[]uint{uint(37340)})&.uint(12):== uint(1009)))|(.uint(14):->contains($uint,[]uint{uint(38121)})&.uint(12):== uint(2000)))|(.uint(14):->contains($uint,[]uint{uint(46270)})&.uint(12):== uint(2022)))|(.uint(14):->contains($uint,[]uint{uint(46432)})&.uint(12):== uint(2111)))|(.uint(14):->contains($uint,[]uint{uint(551496705)})&.uint(12):== uint(10115)))|(.uint(14):->contains($uint,[]uint{uint(60065)})&.uint(12):== uint(10116)))|(.uint(14):->contains($uint,[]uint{uint(15063),uint(19272)})&.uint(12):== uint(10114)))|(.uint(14):->contains($uint,[]uint{uint(22093),uint(22032)})&.uint(12):== uint(10005)))|(.uint(14):->contains($uint,[]uint{uint(26021),uint(26293)})&.uint(12):== uint(7010)))|(.uint(14):->contains($uint,[]uint{uint(29472),uint(29284)})&.uint(12):== uint(10006)))|(.uint(14):->contains($uint,[]uint{uint(30801),uint(32261)})&.uint(12):== uint(1050)))|(.uint(14):->contains($uint,[]uint{uint(32643),uint(32053)})&.uint(12):== uint(3157)))|(.uint(14):->contains($uint,[]uint{uint(33372),uint(33693)})&.uint(12):== uint(3003)))|(.uint(14):->contains($uint,[]uint{uint(48004),uint(38092)})&.uint(12):== uint(2008)))|(.uint(14):->contains($uint,[]uint{uint(12851),uint(13491),uint(15303)})&.uint(12):== uint(40071)))|(.uint(14):->contains($uint,[]uint{uint(12852),uint(11692),uint(15553)})&.uint(12):== uint(50106)))|(.uint(14):->contains($uint,[]uint{uint(15333),uint(11091),uint(14181)})&.uint(12):== uint(50103)))|(.uint(14):->contains($uint,[]uint{uint(16621),uint(12612),uint(16623)})&.uint(12):== uint(40052)))|(.uint(14):->contains($uint,[]uint{uint(21121),uint(21491),uint(21171)})&.uint(12):== uint(10004)))|(.uint(14):->contains($uint,[]uint{uint(34362),uint(34252),uint(34093)})&.uint(12):== uint(40102)))|(.uint(14):->contains($uint,[]uint{uint(11071),uint(12722),uint(14693),uint(11223)})&.uint(12):== uint(10118)))|(.uint(14):->contains($uint,[]uint{uint(13901),uint(12241),uint(15321),uint(19011)})&.uint(12):== uint(40105)))|(.uint(14):->contains($uint,[]uint{uint(16113),uint(10142),uint(15621),uint(14173)})&.uint(12):== uint(50104)))|(.uint(14):->contains($uint,[]uint{uint(16531),uint(11123),uint(19441),uint(15263)})&.uint(12):== uint(4005)))|(.uint(14):->contains($uint,[]uint{uint(16621),uint(15342),uint(16623),uint(11123)})&.uint(12):== uint(40101)))|(.uint(14):->contains($uint,[]uint{uint(24543),uint(23181),uint(24881),uint(23032)})&.uint(12):== uint(20105)))|(.uint(14):->contains($uint,[]uint{uint(34993),uint(31893),uint(34232),uint(34412)})&.uint(12):== uint(40114)))|(.uint(14):->contains($uint,[]uint{uint(46301),uint(46431),uint(46322),uint(46392)})&.uint(12):== uint(20104)))|(.uint(14):->contains($uint,[]uint{uint(46681),uint(46831),uint(34002),uint(34442)})&.uint(12):== uint(3158)))|(.uint(14):->contains($uint,[]uint{uint(11973),uint(11782),uint(11993),uint(12692),uint(11953)})&.uint(12):== uint(8006)))|(.uint(14):->contains($uint,[]uint{uint(26491),uint(24292),uint(26133),uint(20651),uint(26113)})&.uint(12):== uint(7012)))|(.uint(14):->contains($uint,[]uint{uint(10691),uint(11931),uint(14121),uint(14841),uint(13481),uint(14773)})&.uint(12):== uint(50102)))|(.uint(14):->contains($uint,[]uint{uint(11251),uint(12811),uint(14721),uint(14112),uint(11553),uint(11371)})&.uint(12):== uint(3002)))|(.uint(14):->contains($uint,[]uint{uint(16671),uint(13601),uint(12333),uint(14883),uint(14163),uint(12592),uint(11693)})&.uint(12):== uint(4003)))|(.uint(14):->contains($uint,[]uint{uint(32532),uint(32001),uint(32543),uint(30801),uint(46261),uint(32262),uint(32053)})&.uint(12):== uint(2002)))|(.uint(14):->contains($uint,[]uint{uint(11557),uint(11303),uint(11552),uint(14117),uint(14112),uint(14115),uint(64187),uint(64185)})&.uint(12):== uint(10102)))|(.uint(14):->contains($uint,[]uint{uint(12521),uint(16551),uint(11283),uint(11262),uint(16553),uint(15041),uint(11023),uint(15281)})&.uint(12):== uint(18008)))|(.uint(14):->contains($uint,[]uint{uint(17013),uint(13901),uint(13542),uint(54441),uint(51501),uint(55542),uint(12241),uint(16023)})&.uint(12):== uint(4004)))|(.uint(14):->contains($uint,[]uint{uint(34141),uint(34563),uint(34001),uint(34032),uint(46681),uint(34961),uint(46831),uint(34442)})&.uint(12):== uint(2012)))|(.uint(14):->contains($uint,[]uint{uint(15803),uint(12163),uint(14681),uint(11932),uint(11223),uint(14693),uint(12811),uint(11852),uint(11071),uint(11001),uint(14771)})&.uint(12):== uint(3001)))|(.uint(14):->contains($uint,[]uint{uint(11002),uint(14121),uint(15611),uint(17733),uint(13483),uint(14841),uint(12931),uint(12933),uint(11833),uint(12581),uint(14303),uint(14983),uint(13781)})&.uint(12):== uint(3005)))|(.uint(14):->contains($uint,[]uint{uint(24683),uint(24802),uint(24532),uint(24612),uint(24543),uint(24652),uint(23043),uint(23032),uint(23193),uint(27832),uint(23012),uint(24562),uint(24741)})&.uint(12):== uint(7005)))|(.uint(14):->contains($uint,[]uint{uint(46402),uint(34152),uint(46392),uint(46153),uint(46363),uint(46642),uint(46331),uint(46841),uint(46112),uint(31051),uint(46464),uint(46333),uint(46312),uint(46302)})&.uint(12):== uint(2013)))|(.uint(14):->contains($uint,[]uint{uint(14172),uint(11091),uint(14562),uint(14203),uint(11473),uint(15051),uint(15191),uint(12913),uint(14521),uint(15621),uint(14862),uint(16112),uint(14181),uint(13553),uint(16063)})&.uint(12):== uint(4007)))|(.uint(14):->contains($uint,[]uint{uint(31893),uint(34422),uint(34161),uint(34851),uint(34691),uint(34992),uint(34042),uint(34412),uint(34251),uint(34342),uint(34652),uint(34332),uint(34361),uint(34001),uint(34253),uint(34023),uint(34221),uint(34232)})&.uint(12):== uint(2011))))|(.uint(10):== "03"&(((((((((((((((((.uint(14):->contains($uint,[]uint{uint(1)})&.uint(12):== uint(11))|(.uint(14):->contains($uint,[]uint{uint(12542)})&.uint(12):== uint(30261)))|(.uint(14):->contains($uint,[]uint{uint(15753)})&.uint(12):== uint(50161)))|(.uint(14):->contains($uint,[]uint{uint(21573)})&.uint(12):== uint(40119)))|(.uint(14):->contains($uint,[]uint{uint(31102)})&.uint(12):== uint(30211)))|(.uint(14):->contains($uint,[]uint{uint(3262)})&.uint(12):== uint(10164)))|(.uint(14):->contains($uint,[]uint{uint(40522)})&.uint(12):== uint(40111)))|(.uint(14):->contains($uint,[]uint{uint(8491)})&.uint(12):== uint(40161)))|(.uint(14):->contains($uint,[]uint{uint(1241),uint(3203)})&.uint(12):== uint(10161)))|(.uint(14):->contains($uint,[]uint{uint(20821),uint(20703)})&.uint(12):== uint(20164)))|(.uint(14):->contains($uint,[]uint{uint(12251),uint(18083),uint(11051),uint(11112)})&.uint(12):== uint(10112)))|(.uint(14):->contains($uint,[]uint{uint(22072),uint(21042),uint(21472),uint(20383)})&.uint(12):== uint(20114)))|(.uint(14):->contains($uint,[]uint{uint(1091),uint(2212),uint(2742),uint(4283),uint(583)})&.uint(12):== uint(10162)))|(.uint(14):->contains($uint,[]uint{uint(1503),uint(10753),uint(8153),uint(281),uint(7122),uint(813),uint(9921)})&.uint(12):== uint(10163)))|(.uint(14):->contains($uint,[]uint{uint(20241),uint(22292),uint(23841),uint(21732),uint(22463),uint(20892),uint(21263),uint(22342)})&.uint(12):== uint(20116)))|(.uint(14):->contains($uint,[]uint{uint(61291),uint(60063),uint(60042),uint(60031),uint(61251),uint(60092),uint(60421),uint(60243)})&.uint(12):== uint(60111)))|(.uint(14):->contains($uint,[]uint{uint(50321),uint(50673),uint(50252),uint(50203),uint(50562),uint(50242),uint(50323),uint(51221),uint(50791)})&.uint(12):== uint(50111))))))`

	grammar, err := rparser.ParseGrammar(long)
	if err != nil {
		t.Fatal(err)
	}

	parser := katyproto.NewProtoNumParser("testproto", "BananaTuple", testproto.BananaDescription())
	popr := rand.New(rand.NewSource(time.Now().UnixNano()))
	tup := testproto.NewPopulatedBananaTuple(popr, true)
	tup.Peel = uint32(551496705)

	data, err := gogoproto.Marshal(tup)
	if err != nil {
		panic(err)
	}
	parser.Init(data)

	test(t, grammar, parser, true, "Large match test", true)
}
