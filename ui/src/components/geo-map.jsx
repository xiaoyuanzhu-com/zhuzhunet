import { Map, View } from 'ol'
import { Point } from 'ol/geom.js'
import { Feature } from 'ol/index.js'
import { Vector as VectorLayer } from 'ol/layer.js'
import TileLayer from 'ol/layer/Tile'
import "ol/ol.css"
import { useGeographic } from 'ol/proj.js'
import { Vector as VectorSource } from 'ol/source.js'
import OSM from 'ol/source/OSM'
import { useEffect, useRef } from 'react'

useGeographic()

export default function GeoMap({ longitude, latitude }) {
  if (longitude == 0 && latitude == 0) {
    return <div>unknown</div>
  }
  console.log('longitude', longitude, 'latitude', latitude)
  const mapDiv = useRef(null)
  useEffect(() => {
    const point = new Point([longitude, latitude])
    console.log('point', point)
    const map = new Map({
      target: mapDiv.current,
      layers: [
        new TileLayer({ source: new OSM() }),
        new VectorLayer({
          source: new VectorSource({ features: [new Feature(point)] }),
          style: {
            'circle-radius': 6,
            'circle-fill-color': 'red',
          }
        })
      ],
      view: new View({ center: [longitude, latitude], zoom: 4 })
    })
    return () => {
      map.dispose()
    }
  }, [longitude, latitude])
  return (
    <div ref={mapDiv} className="w-full h-full">
    </div>
  )
}
