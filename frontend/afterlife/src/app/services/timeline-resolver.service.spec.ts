import { TestBed } from '@angular/core/testing';

import { TimelineResolverService } from './timeline-resolver.service';

describe('TimelineResolverService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: TimelineResolverService = TestBed.get(TimelineResolverService);
    expect(service).toBeTruthy();
  });
});
