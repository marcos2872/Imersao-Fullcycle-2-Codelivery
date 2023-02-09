import { Prop, Schema, raw, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type RouteDocument = Route & Document;

@Schema()
export class Route {
  @Prop()
  title: string;

  @Prop(
    raw({
      lat: { type: Number },
      long: { type: Number },
    }),
  )
  startPosition: { lat: number; long: number };

  @Prop(
    raw({
      lat: { type: Number },
      long: { type: Number },
    }),
  )
  enPosition: { lat: number; long: number };
}

export const RouteSchema = SchemaFactory.createForClass(Route);
