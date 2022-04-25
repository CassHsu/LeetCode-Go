type CheckInEntry struct {
    StartStation string
    StartTime int
}

type UndergroundSystem struct {
    CheckInMap map[int]CheckInEntry
    JourneyMap map[string][]float64
}


func Constructor() UndergroundSystem {
    checkIn := make(map[int]CheckInEntry)
    journey := make(map[string][]float64)
    
    return UndergroundSystem { CheckInMap: checkIn, JourneyMap: journey }
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    this.CheckInMap[id] = CheckInEntry{ stationName, t }
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
    checkIn := this.CheckInMap[id]
    delete(this.CheckInMap, id)
    
    key := checkIn.StartStation + "-" + stationName
    
    pair := make([]float64, 2)
    if v, ok := this.JourneyMap[key]; !ok {
        this.JourneyMap[key] = pair
    } else {
        pair = v
    }
    
    pair[0] += float64(t - checkIn.StartTime)
    pair[1]++
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    k := startStation + "-" + endStation
    v := this.JourneyMap[k]
    return v[0] / v[1]
}
